package client

import (
	"github.com/laamho/turbo/common"
	"github.com/xtls/xray-core/app/proxyman/command"
	"google.golang.org/grpc"
	"net"
)

func NewServiceClient(ip, port string) command.HandlerServiceClient {
	addr := net.JoinHostPort(ip, port)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	common.Silent(err)
	return command.NewHandlerServiceClient(conn)
}
