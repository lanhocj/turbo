package controller

import (
	"github.com/laamho/turbo/app/service/transport"
	"github.com/laamho/turbo/app/service/transport/client"
	"github.com/laamho/turbo/common"
	"github.com/laamho/turbo/common/orm"
	"log"
	"strings"
	"sync"
)

func deleteNodesByEmail(email string) {
	var nodes []*orm.Node
	var lock sync.Mutex

	if r := orm.DB().Model(&orm.Node{}).Find(&nodes); r.Error != nil {
		log.Println(r.Error.Error())
	}

	for _, node := range nodes {
		go func(node *orm.Node) {
			lock.Lock()
			defer lock.Unlock()

			c := client.NewServiceClient(node.NodeAddr, node.NodePort)
			tag := strings.ToLower(node.NodeTag)

			// 删除用户
			if err := transport.RemoveUser(tag, email, c); err != nil {
				common.Silent(err)
			}
		}(node)
	}

	log.Println("删除节点任务处理完成")
}

func refreshNodesByUser(user *orm.User) {
	email := user.Email
	var lock sync.Mutex

	var allNodes []*orm.Node
	orm.DB().Preload("Nodes").Model(user).Where("email=?", email).First(&user)
	orm.DB().Model(&orm.Node{}).Find(&allNodes)

	var nids []uint

	for _, node := range user.Nodes {
		nids = append(nids, node.ID)
	}

	for _, node := range allNodes {
		go func(node *orm.Node, user *orm.User) {
			lock.Lock()
			defer lock.Unlock()

			c := client.NewServiceClient(node.NodeAddr, node.NodePort)
			tag := strings.ToLower(node.NodeTag)

			// 删除用户
			if err := transport.RemoveUser(tag, user.Email, c); err != nil {
				common.Silent(err)
			} else {
				log.Printf("成功删除：[%s(%s)]->[%s]\n", user.Email, user.Token, node.NodeAddr)
			}

			// 刷新用户
			if common.Combine(node.ID, nids) && user.Role != orm.LAVEL_USER_BLOCK {
				if err := transport.AddUser(tag, user.Email, user.Token, 0, c); err != nil {
					common.Silent(err)
					log.Printf("用户添加失败：[%s(%s)] error: %s \n", user.Email, node.NodeAddr, err.Error())
				} else {
					log.Printf("用户添加失败：[%s(%s)]->[%s]\n", user.Email, user.Token, node.NodeAddr)
				}
			}
		}(node, user)
	}
}
