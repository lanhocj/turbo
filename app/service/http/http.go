package http

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/laamho/turbo/app/controller"
	"github.com/laamho/turbo/app/service/middleware"
	"github.com/laamho/turbo/common/config"
	"github.com/laamho/turbo/common/util"
	"log"
)

func StartWebApplication() error {
	// new gin engine..
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.LoadHTMLFiles(util.RelativeFilePath("template/views")...)
	r.Static("/static", "template/static")

	store := cookie.NewStore(config.GetSecret())
	r.Use(sessions.Sessions("turbo", store))
	addr := config.GetAddr()

	controller.GlobalData.Init()

	front := r.Group("/")
	{
		front.GET("/login", controller.LoginHandler())
		front.GET("/register")
		front.GET("/c/:id")

		front.Use(middleware.Authenticator())
		{
			front.GET("/")
			front.GET("/nodes")
			front.GET("/users")
			front.GET("/config")
		}
	}

	log.Printf("Listen at %s\n", addr)
	return r.Run(addr)
}
