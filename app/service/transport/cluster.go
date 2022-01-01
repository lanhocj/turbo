package transport

import (
	"context"
	"github.com/xtls/xray-core/app/proxyman"
	"github.com/xtls/xray-core/app/proxyman/command"
	"github.com/xtls/xray-core/common/net"
	"github.com/xtls/xray-core/common/serial"
	core "github.com/xtls/xray-core/core"
	"github.com/xtls/xray-core/proxy/trojan"
	"github.com/xtls/xray-core/transport/internet"
	"strings"
)

// AddInboundProxy 存在bug...
func AddInboundProxy(tag string, addr string, port net.Port, c command.HandlerServiceClient) (*command.AddInboundResponse, error) {
	tag = strings.ToUpper(tag)
	portRange := net.SinglePortRange(port)

	return c.AddInbound(context.Background(), &command.AddInboundRequest{
		Inbound: &core.InboundHandlerConfig{
			Tag: tag,
			ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
				PortList: &net.PortList{
					Range: []*net.PortRange{portRange},
				},
				//Listen: net.NewIPOrDomain(localAddr),
				AllocationStrategy: &proxyman.AllocationStrategy{
					Type: proxyman.AllocationStrategy_Always,
					Concurrency: &proxyman.AllocationStrategy_AllocationStrategyConcurrency{
						Value: 1,
					},
				},
				StreamSettings: &internet.StreamConfig{
					ProtocolName: "tcp",
				},
				ReceiveOriginalDestination: false,
				SniffingSettings: &proxyman.SniffingConfig{
					Enabled:             true,
					DestinationOverride: []string{"http", "tls"},
				},
			}),
			ProxySettings: serial.ToTypedMessage(&trojan.ServerConfig{}),
		},
	})
}
