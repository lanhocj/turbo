package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authenticator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		token := session.Get("token")

		if token == nil {
			ctx.Redirect(302, "/login")
			return
		}

		ctx.Next()
	}
}
