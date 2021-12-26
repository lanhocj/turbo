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

type Token struct {
	gorm.Model

	Token  string
	UserID int
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type User struct {
	gorm.Model

	Email    string `json:"email"`
	Password string `json:"password"`
	Role     int    `json:"role"`
	Hash     string `json:"hash"`

	Nodes []Node `gorm:"many2many:node_user;"`
}
