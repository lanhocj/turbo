package client

import (
	"context"
	"fmt"
	"github.com/laamho/turbo/common"
	"github.com/xtls/xray-core/app/proxyman/command"
	"github.com/xtls/xray-core/common/serial"
	"google.golang.org/grpc"
	"net"
	"regexp"
	"strings"
)

type NodeState struct {
	Code string
	Desc string
}

func (s *NodeState) State() {}

// rpc error: code = Unknown desc = app/proxyman/command: failed to get handler: PROXY > app/proxyman/inbound: handler not found: PROXY
// rpc error: code = Unavailable desc = connection error: desc = "transport: Error while dialing dial tcp 127.0.0.1:3333: connect: connection refused"
func NewState(s string) {
	re := regexp.MustCompile("code = (.+?) desc = (.+?):")
	r := re.FindString(s)
	fmt.Println(r)
}

func NewServiceClient(ip, port string) command.HandlerServiceClient {
	addr := net.JoinHostPort(ip, port)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	common.Silent(err)
	return command.NewHandlerServiceClient(conn)
}

func NewTestServiceClient(ip, port, tag string) bool {
	c := NewServiceClient(ip, port)
	tag = strings.ToUpper(tag)
	_, err := c.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag:       tag,
		Operation: serial.ToTypedMessage(&command.AddUserOperation{}),
	})
	// 如果 连接拒绝 返回 true..
	ok, _ := regexp.MatchString(`connection refused`, err.Error())
	return ok
}
