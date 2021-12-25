package http

import (
	"github.com/gin-gonic/gin"
	"github.com/laamho/turbo/common/config"
)

// Start a server
func Start() error {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Static("/static", "./static")

	addr := config.GetAddr()

	r.LoadHTMLGlob("template/*.tmpl")

	api := r.Group("/api")
	{
		api.GET("/register", ApiRegisterUserHandler)

		api.Use()
		{
			api.GET("/nodes")          // 列出所有节点
			api.GET("/:user_id/nodes") // 通过用户列出节点
		}
	}

	front := r.Group("/")
	{
		front.GET("/", IndexHandler)
		front.GET("/nodes")
		front.GET("/users", ListUserHandler)
		front.GET("/config")
	}

	return r.Run(addr)
}
