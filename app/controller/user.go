package controller

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/laamho/turbo/app/controller/internal"
	"github.com/laamho/turbo/app/service/clash"
	"github.com/laamho/turbo/common"
	"github.com/laamho/turbo/common/orm"
	"github.com/laamho/turbo/common/util"
	"github.com/xtls/xray-core/common/uuid"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func UserListHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []orm.User
		//var nodes []orm.Node
		orm.DB().Model(&orm.User{}).Preload("Nodes").Find(&users)
		var u *orm.User

		curr, ok := c.Get("currentUser")
		if ok {
			u = curr.(*orm.User)
		}

		var r internal.UserListResponse

		for _, user := range users {
			if 1 != u.ID && user.ID == 1 {
				continue
			}

			var un internal.NodeListResponse
			var role string
			switch user.Role {
			case orm.LEVEL_USER_ADMIN:
				role = "管理员"
				break
			case orm.LAVEL_USER_BLOCK:
				role = "禁止用户"
				break
			case orm.LAVEL_USER_NORMAL:
				role = "普通用户"
				break
			default:
				role = "禁止用户"
				break
			}

			u := &internal.UserResponse{
				Email:   user.Email,
				RoleId:  user.Role,
				Locked:  user.Locked,
				IsAdmin: orm.LEVEL_USER_ADMIN == user.Role,
				Role:    role,
				NodeNum: len(user.Nodes),
				Token:   user.Token,
			}

			for _, node := range user.Nodes {
				un = append(un, &internal.NodeResponse{
					Addr:   node.ClientAddr,
					Region: "不知道",
				})
			}

			u.NodeList = un
			r = append(r, u)
		}

		c.JSON(200, r)
		return
	}
}

func UserAddHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestUser := internal.UserCreateRequest{}

		if err := c.ShouldBindJSON(&requestUser); err != nil {
			c.AbortWithStatusJSON(416, gin.H{"message": "表单提交错误", "error": err.Error()})
			return
		}

		user := new(orm.User)

		if r := orm.DB().Model(&user).Where("email=?", requestUser.Email).First(&user); errors.Is(r.Error, gorm.ErrRecordNotFound) {
			user.Email = requestUser.Email
			user.Password = util.Hash(user.Email, requestUser.Password)
			token := uuid.New()
			user.Token = token.String()
			orm.DB().Model(&user).Create(&user)

			c.JSON(200, user)
			return
		}

		internal.Error(c, 20001, gin.H{"message": fmt.Sprintf("账户 [%s] 已存在！", requestUser.Email)})
		return
	}
}

func UserNodesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqeustData := &internal.RequestWithEmail{}
		if err := c.ShouldBindJSON(&reqeustData); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"message": "请求错误", "error": err.Error()})
			return
		}

		var nodes []orm.Node
		orm.DB().Model(&orm.Node{}).Order("node_name ASC").Find(&nodes)
		var user orm.User

		if r := orm.DB().Model(&user).Preload("Nodes").Where("email=?", reqeustData.Email).First(&user); errors.Is(r.Error, gorm.ErrRecordNotFound) {
			common.Silent(r.Error)
		}

		var r internal.UserNodesResponse

		var curUserNodes []uint

		for _, cur := range user.Nodes {
			curUserNodes = append(curUserNodes, cur.ID)
		}

		fmt.Println(curUserNodes)

		for _, node := range nodes {
			n := &internal.UserNodeResponse{
				Name:  node.NodeName,
				ID:    node.ID,
				Using: common.Combine(node.ID, curUserNodes),
			}
			r = append(r, n)
		}

		c.JSON(200, r)
		return
	}
}

func LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()

		if err := session.Save(); err != nil {
			c.AbortWithStatus(500)
			return
		}
		c.Redirect(302, "/login")
	}
}

func UserSettingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData = internal.PutUserToNodeRequest{}

		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.AbortWithStatusJSON(200, gin.H{"data": requestData, "error": err.Error()})
			return
		}

		user := &orm.User{}
		orm.DB().Model(&user).Where("email=?", requestData.Email).First(&user)
		nodeIds := strings.Split(requestData.Node, ",")
		var nids []uint
		for _, n := range nodeIds {
			c, _ := strconv.Atoi(n)
			nids = append(nids, uint(c))
		}

		orm.DB().Model(&user).Association("Nodes").Clear()
		var nodes []orm.Node
		orm.DB().Model(&orm.Node{}).Find(&nodes, nids)
		user.Nodes = nodes

		role, _ := strconv.Atoi(requestData.Role)
		user.Role = role

		go refreshNodesByUser(user)

		if r := orm.DB().Save(&user); r.Error == nil {
			c.JSON(200, gin.H{"status": 2000, "message": "数据添加成功"})
			return
		}

		c.JSON(200, gin.H{"message": "数据添加失败"})
		return
	}
}

func UserRemoveHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := new(internal.RequestWithEmail)

		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(200, gin.H{"message": "更新失败", "error": err.Error()})
			return
		}

		user := new(orm.User)
		if err := orm.DB().Model(&user).Where("email=?", request.Email).Association("Nodes").Clear(); err != nil {
			common.Silent(err)
		}

		go deleteNodesByEmail(request.Email)

		if r := orm.DB().Model(user).Where("email=?", request.Email).Delete(&user); errors.Is(r.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(200, gin.H{"message": "没有找到这个账户", "error": r.Error})
			return
		}

		c.AsciiJSON(200, gin.H{"message": "账户删除成功"})
		return
	}
}

func UserPasswordChangeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := new(internal.RequestWithEmailAndPassword)
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(200, gin.H{"data": request, "error": err.Error()})
			return
		}

		var user orm.User
		if err := orm.DB().Model(&user).Where("email=?", request.Email).First(&user); errors.Is(err.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(200, gin.H{"data": request, "error": err.Error})
			return
		}

		user.Password = util.Hash(user.Email, request.Password)
		if r := orm.DB().Save(&user); r.Error == nil {
			c.JSON(200, gin.H{"status": 2000, "message": "数据添加成功"})
			return
		}

		c.JSON(200, gin.H{"message": "数据添加失败"})
	}
}

func GetUserConfigPath() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Param("token")
		user := new(orm.User)

		if r := orm.DB().Preload("Nodes", func(db *gorm.DB) *gorm.DB { return db.Order("node_name ASC") }).Where("token=?", token).First(&user); errors.Is(r.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(404)
			return
		}

		go refreshNodesByUser(user)

		obj := clash.Default()

		for _, node := range user.Nodes {
			obj.AddProxy("trojan", node.NodeName, node.ClientAddr, node.ClientPort, user.Token, false)
		}

		c.String(200, obj.String())
		return
	}
}
