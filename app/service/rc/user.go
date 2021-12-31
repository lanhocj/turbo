package rc

import (
	"context"
	"github.com/xtls/xray-core/app/proxyman/command"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/proxy/trojan"
	"strings"
)

// AddUser to cluster through xray api
func AddUser(tag, email, password string, level uint32, c command.HandlerServiceClient) error {
	tag = strings.ToLower(tag)

	_, err := c.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: tag,
		Operation: serial.ToTypedMessage(&command.AddUserOperation{
			User: &protocol.User{
				Email: email,
				Account: serial.ToTypedMessage(&trojan.Account{
					Password: password,
					Flow:     trojan.XRD,
				}),
			},
		}),
	})

	if err != nil {
		return err
	}
	return nil
}

// RemoveUser to cluster through xray api
func RemoveUser(tag, email string, c command.HandlerServiceClient) error {
	tag = strings.ToLower(tag)

	_, err := c.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: tag,
		Operation: serial.ToTypedMessage(&command.RemoveUserOperation{
			Email: email,
		}),
	})

	return err
}
