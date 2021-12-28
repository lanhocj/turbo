package lpc

import (
	"encoding/json"
	"fmt"
	"github.com/laamho/turbo/common"
	"github.com/laamho/turbo/common/orm"
	"github.com/laamho/turbo/common/util"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type RpcService struct{}

type RPC struct {
	client *rpc.Client
}

const (
	CreateUserService = "Rpc.CreateUser"
	RemoveUserService = "Rpc.RemoveUser"
)

func (p *RPC) Listen() error {
	if err := rpc.RegisterName("Rpc", new(RpcService)); err != nil {
		return err
	}

	ln, _ := net.ListenTCP("tcp", &net.TCPAddr{
		IP: net.IPv4(127, 0, 0, 1), Port: 9908,
	})

	for {
		conn, _ := ln.Accept()

		go func() {
			err := rpc.ServeRequest(jsonrpc.NewServerCodec(conn))

			if err != nil {
				log.Println(err.Error())
			}
		}()
	}
}

func (p *RPC) Dial() error {
	conn, err := net.Dial("tcp", "localhost:9908")
	if err != nil {
		return err
	}

	p.client = rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return nil
}

func (p *RPC) Send(name string, action func() interface{}) (string, error) {
	var reply string
	data := action()
	err := p.client.Call(name, data, &reply)
	return reply, err
}

func (r *RpcService) CreateUser(req interface{}, res *string) error {
	user := orm.User{}

	enc, err := json.Marshal(req)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(enc, &user); err != nil {
		return err
	}

	user.Password = util.Hash(user.Email, user.Password)

	log.Printf("Create user via command line [%s]", user.Email)
	orm.DB().Create(&user)

	reply := fmt.Sprintf("\033[2mAccount [%s] already created.\033[0m\n", user.Email)
	res = &reply

	return nil
}

func (r *RpcService) RemoveUser(req interface{}, res *string) error {
	user := orm.User{}

	enc, _ := json.Marshal(req)

	common.Must(json.Unmarshal(enc, &user))

	orm.DB().Unscoped().Delete(&user, "email=?", user.Email)

	reply := fmt.Sprintf("\033[2mAccount [%s] already deleted.\033[0m\n", user.Email)
	res = &reply

	return nil
}
