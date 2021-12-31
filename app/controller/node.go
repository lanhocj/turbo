package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/laamho/turbo/app/controller/internal"
	"github.com/laamho/turbo/app/service/rc/client"
)

func NodeAvailableTestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data internal.NodeAvailableRequest

		if err := c.ShouldBindJSON(&data); err != nil {
			internal.Error(c, 20001, err.Error())
			return
		}

		res := new(internal.NodeAvailableResponse)

		if s := client.NewTestServiceClient(data.Addr, data.Port, data.Tag); s {
			res.State = 0
		} else {
			res.State = 1
		}

		res.Build(c)
		return
	}
}
