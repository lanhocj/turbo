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

	addr := config.GetAddr()

	r.LoadHTMLGlob("template/*.tmpl")

	api := r.Group("/api")
	{
		api.GET("/")
		api.GET("/register")

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
		front.GET("/users")
		front.GET("/config")
	}

	return r.Run(addr)
}
