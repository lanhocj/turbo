package orm

import (
	"gorm.io/gorm"
	"time"
)

const (
	LEVEL_USER_ADMIN int = iota
	LAVEL_USER_NORMAL
	LAVEL_USER_BLOCK
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
	Password string
	Role     int `json:"role"`

	Nodes []Node `gorm:"many2many:node_user;"`
}
