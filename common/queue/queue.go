package queue

import (
	"github.com/laamho/turbo/app/service/rc"
	"github.com/laamho/turbo/app/service/rc/client"
	"github.com/laamho/turbo/common"
	"github.com/laamho/turbo/common/orm"
	"time"
)

func loader() {
	var nodes []orm.Node
	orm.DB().Preload("Users").Find(&nodes)

	for _, node := range nodes {
		c := client.NewServiceClient(node.NodeAddr, node.NodePort)
		// 如果 连接拒绝 返回 true..

		for _, user := range node.Users {
			if user.Role != orm.LAVEL_USER_BLOCK {
				if err := rc.AddUser(node.NodeTag, user.Email, user.Token, 1, c); err != nil {
					common.Silent(err)
				}
			}
		}
	}

	time.AfterFunc(time.Minute*time.Duration(10), loader)
}

func SyncWithNodes() {
	go loader()
}
