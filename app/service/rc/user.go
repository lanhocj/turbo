package rc

import (
	"context"
	"github.com/xtls/xray-core/app/proxyman/command"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/proxy/trojan"
	"log"
)

// AddUser to cluster through xray api
func AddUser(email, password string, level uint32, c command.HandlerServiceClient) error {
	resp, err := c.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: "proxy",
		Operation: serial.ToTypedMessage(&command.AddUserOperation{
			User: &protocol.User{
				Email: email,
				Level: level,
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

	log.Printf("Ok, %s", resp.String())
	return nil
}

// RemoveUser to cluster through xray api
func RemoveUser(email string, c command.HandlerServiceClient) error {
	resp, err := c.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: "proxy",
		Operation: serial.ToTypedMessage(&command.RemoveUserOperation{
			Email: email,
		}),
	})

	if err != nil {
		return err
	}

	log.Printf("Ok, %s", resp.String())
	return nil
}
