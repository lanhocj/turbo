package orm

import (
	"context"
	"github.com/xtls/xray-core/common/uuid"
	"gorm.io/gorm"
)

const (
	LAVEL_USER_BLOCK int = iota
	LEVEL_USER_ADMIN
	LAVEL_USER_NORMAL
)

const (
	ClientProxyMode_Trojan = "trojan"
)

type Node struct {
	gorm.Model

	NodeName string `json:"nodeName"`
	NodeAddr string `json:"nodeAddr"`
	NodePort string `json:"nodePort"`
	NodeTag  string `json:"nodeTag"`

	ClientAddr      string `json:"clientAddr"`
	ClientPort      string `json:"clientPort"`
	ClientProxyMode string `json:"clientProxyMode"`

	Users []User `gorm:"many2many:node_user;"`
}

func (n *Node) Sync(c context.Context) {

}

type Token struct {
	gorm.Model

	Token  string
	UserID int
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type User struct {
	gorm.Model

	Email    string `json:"email" gorm:"email,unique,omitempty"`
	Password string `json:"password" gorm:"password,omitempty"`
	Role     int    `json:"role" gorm:"role,omitempty"`
	Hash     string `json:"-" gorm:"hash,omitempty"`
	Token    string `json:"token"`
	Locked   bool   `json:"locked"`

	Nodes []Node `json:"nodes" gorm:"many2many:node_user;"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	r := uuid.New()
	u.Token = r.String()
	return nil
}
