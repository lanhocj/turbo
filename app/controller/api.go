package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/laamho/turbo/app/controller/internal"
	"github.com/laamho/turbo/app/service/rc"
	"github.com/laamho/turbo/common"
	"github.com/laamho/turbo/common/orm"
	"github.com/xtls/xray-core/app/proxyman/command"
	"google.golang.org/grpc"
	"net"
)

func AddNodeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		node := new(orm.Node)

		if err := c.ShouldBindJSON(&node); err != nil {
			internal.Error(c, 20001, gin.H{"message": "Invalid form", "form": node, "error": err.Error()})
			return
		}

		orm.DB().Create(&node)

		c.JSON(200, node)
		return
	}
}

// AddProxyHandler error.
func AddProxyHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		srv := net.JoinHostPort("127.0.0.1", "3333")
		conn, err := grpc.Dial(srv, grpc.WithInsecure())
		if err != nil {
			context.AbortWithError(500, err)
		}
		client := command.NewHandlerServiceClient(conn)

		res, err := rc.AddInboundProxy("trojan", "0.0.0.0", 10881, client)
		common.Silent(err)

		context.String(200, res.String())
	}
}
