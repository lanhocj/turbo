package controller

import (
	"github.com/gin-gonic/gin"
)

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "login.tmpl.html", gin.H{
			"globals": GlobalData,
		})
	}
}
