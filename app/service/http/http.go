package http

import "github.com/gin-gonic/gin"

func StartWebApplication() error {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	return r.Run()
}
