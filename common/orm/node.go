package orm

import (
	"gorm.io/gorm"
	"time"
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

	UpdatedAt time.Time
	CreatedAt time.Time
}

type User struct {
	gorm.Model

	Email    string `json:"email"`
	Password string `json:"password"`
	Role     int    `json:"role"`
	Token    string `json:"token"`

	Nodes []Node `gorm:"many2many:node_user;"`
}
