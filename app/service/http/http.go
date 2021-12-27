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
		front.POST("/login", controller.RequestLoginHandler())
		front.GET("/register")
		front.GET("/c/:id")

		front.Use(middleware.Authenticator())
		{
			front.GET("/", controller.IndexHandler())
			front.GET("/nodes", controller.NodesListHandler())
			front.GET("/users", controller.UsersListHandler())
			front.GET("/config")
		}
	}

	api := r.Group("api")
	{
		api.GET("/node", controller.AddNodeHandler())
	}

	log.Printf("Listen at %s\n", addr)
	return r.Run(addr)
}
