package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/laamho/turbo/common/util"
)

func Authenticator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		token := session.Get("token")

		if util.Empty(token) {
			ctx.Redirect(302, "/login")
			return
		}

		ctx.Set("token", token)

		ctx.Next()
	}
}
