package controller

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/laamho/turbo/app/controller/internal"
	"github.com/laamho/turbo/common/orm"
	"github.com/laamho/turbo/common/util"
	"github.com/xtls/xray-core/common/uuid"
	"gorm.io/gorm"
)

var globals *internal.GlobalPageData

func init() {
	globals = internal.Glob.Init()
}

func LoginViewHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "login.tmpl.html", gin.H{
			"globals": globals,
		})
	}
}

func RequestLoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		request := new(internal.RequestWithEmailAndPassword)

		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "表单错误"})
			return
		}

		user := new(orm.User)

		if obj := orm.DB().Where("email = ?", request.Email).Find(&user); errors.Is(obj.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(200, gin.H{"message": "账户登录失败，请核对登录邮箱或密码"})
			return
		}

		if util.Valid(request.Email, request.Password, user.Password) {
			hash := util.SecretHash(request.Email)
			user.Hash = hash

			if obj := orm.DB().Model(&user).Where("email=?", request.Email).Update("hash", hash); obj.Error != nil {
				fmt.Println(obj)
				c.AbortWithStatusJSON(200, gin.H{"message": "登录失败，请稍后再试"})
				return
			}

			session.Set("token", hash)

			if err := session.Save(); err != nil {
				c.AbortWithStatusJSON(200, gin.H{"message": "账户登录失败，请核对登录邮箱或密码"})
				return
			}

			c.Redirect(302, "/")
			return
		}

		c.AbortWithStatusJSON(200, gin.H{"message": "账户登录失败，请核对登录邮箱或密码"})
		return
	}
}

func IndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetString("token")
		user := new(orm.User)
		orm.DB().Where("hash", token).Find(&user)

		scheme := "http"
		if c.Request.TLS != nil {
			scheme = "https"
		}

		globals.SetCurrentPath(c.FullPath())
		globals.SetCurrentTitle("配置文件")

		c.HTML(200, "index.tmpl.html", gin.H{
			"globals": globals,
			"User":    user,
			"Url":     fmt.Sprintf("%s://%s/c/%s", scheme, c.Request.Host, user.Token),
		})
		return
	}
}

func UsersListHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetString("token")
		user := new(orm.User)
		orm.DB().Where("hash", token).Find(&user)

		globals.SetCurrentPath(c.FullPath())
		globals.SetCurrentTitle("用户管理")

		c.HTML(200, "users.tmpl.html", gin.H{
			"globals": globals,
			"User":    user,
		})
	}
}

func NodesListHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetString("token")
		user := new(orm.User)
		node := new(orm.Node)
		var nodes []orm.Node
		orm.DB().Where("hash", token).Find(&user)
		orm.DB().Model(&node).Order("node_name ASC").Limit(10).Find(&nodes)

		var nodeResult []map[string]interface{}

		for _, n := range nodes {
			nodeResult = append(nodeResult, map[string]interface{}{
				"addr":  n.NodeAddr,
				"name":  n.NodeName,
				"tag":   n.NodeTag,
				"port":  n.NodePort,
				"state": -1,
			})
		}

		globals.SetCurrentPath(c.FullPath())
		globals.SetCurrentTitle("节点管理")

		c.HTML(200, "nodes.tmpl.html", gin.H{
			"globals": globals,
			"User":    user,
			"Nodes":   nodeResult,
		})
	}
}

func UserLockHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := new(internal.RequestWithEmail)

		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(200, gin.H{"message": "请求错误", "error": err.Error()})
			return
		}

		user := new(orm.User)
		if r := orm.DB().Model(user).Where("email=?", request.Email).First(&user); r.Error != nil {
			c.AbortWithStatusJSON(200, gin.H{"message": "查询错误", "error": r.Error})
			return
		}
		user.Locked = !user.Locked

		go deleteNodesByEmail(request.Email)

		if r := orm.DB().Save(&user); r.Error != nil {
			c.AbortWithStatusJSON(200, gin.H{"message": "操作失败", "error": r.Error})
			return
		}

		c.AsciiJSON(200, gin.H{"message": "成功"})
		return
	}
}

func UserTokenRefreshHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := new(internal.RequestWithEmail)
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(200, gin.H{"message": "请求错误", "error": err.Error()})
			return
		}

		user := new(orm.User)
		if r := orm.DB().Model(user).Where("email=?", request.Email).First(&user); r.Error != nil {
			c.AbortWithStatusJSON(200, gin.H{"message": "查询错误", "error": r.Error})
			return
		}
		token := uuid.New()
		user.Token = token.String()

		if r := orm.DB().Save(&user); r.Error != nil {
			c.AbortWithStatusJSON(200, gin.H{"message": "操作失败", "error": r.Error})
			return
		}

		go refreshNodesByUser(user)

		scheme := "http"
		if c.Request.TLS != nil {
			scheme = "https"
		}

		url := fmt.Sprintf("%s://%s/c/%s", scheme, c.Request.Host, user.Token)
		c.AsciiJSON(200, gin.H{"message": "成功", "url": url})
		return
	}
}
