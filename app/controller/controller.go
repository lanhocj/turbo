package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/laamho/turbo/app/service/rc/client"
	"github.com/laamho/turbo/common/orm"
	"github.com/laamho/turbo/common/util"
	"github.com/xtls/xray-core/common/errors"
)

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "login.tmpl.html", gin.H{
			"globals": GlobalData,
		})
	}
}

func RequestLoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		var email, password string
		var ok bool
		if email, ok = c.GetPostForm("email"); !ok {
			c.AbortWithError(401, errors.New("请输入登录邮箱"))
		}

		if password, ok = c.GetPostForm("password"); !ok {
			c.AbortWithError(401, errors.New("请输入登录密码"))
		}

		user := new(orm.User)

		orm.DB().Where("email = ?", email).Find(&user)
		fmt.Printf("获取到的 orm.User 数据: \n")
		fmt.Println(user)
		fmt.Printf("Id: %d, Email: [%s] and Password [%s]\n", user.ID, user.Email, user.Password)

		if util.Vaild(email, password, user.Password) {
			hash := util.SecretHash(email)
			user.Hash = hash
			result := orm.DB().Model(&user).Where("email=?", email).Update("hash", hash)

			if result.Error != nil {
				c.AbortWithError(500, result.Error)
				return
			}

			session.Set("token", hash)
			if err := session.Save(); err != nil {
				c.AbortWithError(500, err)
			}

			c.Redirect(302, "/")
			return
		}

		c.String(200, fmt.Sprintf("登录账户 [%s] 失败", email))
	}
}

func IndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetString("token")
		user := new(orm.User)
		orm.DB().Where("hash", token).Find(&user)

		c.HTML(200, "index.tmpl.html", gin.H{
			"globals":          GlobalData,
			"CurrentPageTitle": "配置文件",
			"CurrentPath":      c.FullPath(),
			"User":             user,
		})
		return
	}
}

func UsersListHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetString("token")
		user := new(orm.User)
		orm.DB().Where("hash", token).Find(&user)

		c.HTML(200, "users.tmpl.html", gin.H{
			"globals":          GlobalData,
			"CurrentPageTitle": "配置文件",
			"CurrentPath":      c.FullPath(),
			"User":             user,
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
		orm.DB().Model(node).Find(&nodes)

		var nodeResult []map[string]interface{}

		for _, n := range nodes {

			ok := client.NewTestServiceClient(n.NodeAddr, n.NodePort, n.NodeTag)

			nodeResult = append(nodeResult, map[string]interface{}{
				"addr":  n.NodeAddr,
				"name":  n.NodeName,
				"state": !ok,
			})
		}

		fmt.Println(nodeResult)

		c.HTML(200, "nodes.tmpl.html", gin.H{
			"globals":          GlobalData,
			"CurrentPageTitle": "配置文件",
			"CurrentPath":      c.FullPath(),
			"User":             user,
			"Nodes":            nodeResult,
		})
	}
}
