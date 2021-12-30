package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/laamho/turbo/app/controller/structs"
	"github.com/laamho/turbo/common/orm"
	"github.com/laamho/turbo/common/util"
	"github.com/xtls/xray-core/common/errors"
	"github.com/xtls/xray-core/common/uuid"
)

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		GlobalData.SetCurrentPath(c.FullPath())
		GlobalData.SetCurrentTitle("登录")

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

		if util.Valid(email, password, user.Password) {
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

		scheme := "http"
		if c.Request.TLS != nil {
			scheme = "https"
		}

		GlobalData.SetCurrentPath(c.FullPath())
		GlobalData.SetCurrentTitle("配置文件")

		c.HTML(200, "index.tmpl.html", gin.H{
			"globals": GlobalData,
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

		GlobalData.SetCurrentPath(c.FullPath())
		GlobalData.SetCurrentTitle("用户管理")

		c.HTML(200, "users.tmpl.html", gin.H{
			"globals": GlobalData,
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
		orm.DB().Model(node).Find(&nodes)

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

		GlobalData.SetCurrentPath(c.FullPath())
		GlobalData.SetCurrentTitle("节点管理")

		c.HTML(200, "nodes.tmpl.html", gin.H{
			"globals": GlobalData,
			"User":    user,
			"Nodes":   nodeResult,
		})
	}
}

func SetUserLockHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request structs.GetNodeListRequest
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

		if !user.Locked {
			go flushNodesByUser(user)
		}

		if r := orm.DB().Save(&user); r.Error != nil {
			c.AbortWithStatusJSON(200, gin.H{"message": "操作失败", "error": r.Error})
			return
		}

		c.AsciiJSON(200, gin.H{"message": "成功"})
		return
	}
}

func FlushTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request structs.GetNodeListRequest
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

		go flushNodesByUser(user)

		if r := orm.DB().Save(&user); r.Error != nil {
			c.AbortWithStatusJSON(200, gin.H{"message": "操作失败", "error": r.Error})
			return
		}

		scheme := "http"
		if c.Request.TLS != nil {
			scheme = "https"
		}

		url := fmt.Sprintf("%s://%s/c/%s", scheme, c.Request.Host, user.Token)
		c.AsciiJSON(200, gin.H{"message": "成功", "url": url})
		return
	}
}
