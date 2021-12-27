package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/laamho/turbo/app/service/rc"
	"github.com/laamho/turbo/common"
	"github.com/xtls/xray-core/app/proxyman/command"
	"google.golang.org/grpc"
	"net"
)

func AddNodeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//client := r.NewServiceClient()
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

		res, err := rc.AddInboundProxy("trojan", "0.0.0.0", 8881, client)
		common.Silent(err)

		context.String(200, res.String())
	}
}