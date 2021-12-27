package rc

import (
	"context"
	"fmt"
	"github.com/xtls/xray-core/app/proxyman"
	"github.com/xtls/xray-core/app/proxyman/command"
	"github.com/xtls/xray-core/common/net"
	"github.com/xtls/xray-core/common/serial"
	core "github.com/xtls/xray-core/core"
	"github.com/xtls/xray-core/proxy/trojan"
	"strings"
)

// AddInboundProxy 存在bug...
func AddInboundProxy(tag string, addr string, port net.Port, c command.HandlerServiceClient) (*command.AddInboundResponse, error) {
	tag = strings.ToUpper(tag)
	portRange := net.SinglePortRange(port)
	localAddr := net.ParseAddress(addr)
	fmt.Println(net.NewIPOrDomain(localAddr).GetAddress())

	return c.AddInbound(context.Background(), &command.AddInboundRequest{
		Inbound: &core.InboundHandlerConfig{
			Tag: tag,
			ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
				PortList: &net.PortList{
					Range: []*net.PortRange{portRange},
				},
				Listen: net.NewIPOrDomain(localAddr),
			}),
			ProxySettings: serial.ToTypedMessage(&trojan.ServerConfig{}),
		},
	})
}
