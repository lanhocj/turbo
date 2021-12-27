package orm

import (
	"gorm.io/gorm"
)

const (
	LAVEL_USER_BLOCK int = iota
	LEVEL_USER_ADMIN
	LAVEL_USER_NORMAL
)

type Node struct {
	gorm.Model

	Addr, Port, Tag string

	Users []User `gorm:"many2many:node_user;"`
}

type Token struct {
	gorm.Model

	Token  string
	UserID int
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type User struct {
	gorm.Model

	Email    string `json:"email" gorm:"email,omitempty"`
	Password string `json:"-" gorm:"password,omitempty"`
	Role     int    `json:"role" gorm:"role,omitempty"`
	Hash     string `json:"-" gorm:"hash,omitempty"`

	Nodes []Node `json:"nodes" gorm:"many2many:node_user;"`
}
