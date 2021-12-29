package middleware

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/laamho/turbo/common/orm"
	"github.com/laamho/turbo/common/util"
	"gorm.io/gorm"
)

func Authenticator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := &orm.User{}
		session := sessions.Default(ctx)
		token := session.Get("token")
		r := orm.DB().Model(user).Where("hash=?", token)

		if util.Empty(token) || errors.Is(r.Error, gorm.ErrRecordNotFound) {
			ctx.Redirect(302, "/login")
			return
		}

		ctx.Set("token", token)

		ctx.Next()
	}
}
