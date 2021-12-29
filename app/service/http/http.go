package http

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/laamho/turbo/app/controller"
	"github.com/laamho/turbo/app/service/middleware"
	"github.com/laamho/turbo/common/config"
	"github.com/laamho/turbo/common/queue"
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
		front.GET("/c/:token", controller.GetUserConfigPath())

		front.Use(middleware.Authenticator())
		{
			front.GET("/", controller.IndexHandler())
			front.GET("/nodes", controller.NodesListHandler())
			front.GET("/users", controller.UsersListHandler())
			front.GET("/config")
		}
	}

	api := r.Group("api", middleware.Authenticator())
	{
		node := api.Group("/node")
		{
			node.POST("", controller.AddNodeHandler())
			node.POST("/available", controller.NodeAvailableTestHandler())
		}

		user := api.Group("users")
		{
			user.GET("/", controller.UserListHandler())
			user.POST("/create", controller.AddUserHandler())
			user.POST("/nodes", controller.GetNodeWithUser())
			user.POST("/nodeSetting", controller.PutUserToNode())
			user.POST("/changeUserPassword", controller.PutChangeUserPassword())
		}

		//api.GET("/createNode", controller.AddProxyHandler()) // 添加代理失败，原因未知。
	}

	queue.SyncWithNodes()

	log.Printf("Listen at %s\n", addr)
	return r.Run(addr)
}
