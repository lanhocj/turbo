package http

import (
	"github.com/gin-gonic/gin"
	"github.com/laamho/turbo/common/orm"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func ListUserHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "users.tmpl", gin.H{})
}

func ApiRegisterUserHandler(c *gin.Context) {
	var user orm.User
	c.Bind(&user)
}
