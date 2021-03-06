package orm

import (
	"github.com/laamho/turbo/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
)

type Object struct {
	*gorm.DB

	Driver, DSN string
}

var db *gorm.DB

func (o *Object) open() (*gorm.DB, error) {
	switch strings.ToLower(o.Driver) {
	case "mysql":
		return gorm.Open(mysql.Open(o.DSN))
	case "sqlite":
		return gorm.Open(sqlite.Open(o.DSN))
	default:
		return gorm.Open(sqlite.Open(o.DSN))
	}
}

func AutoMigrate(dst ...interface{}) error {
	return db.AutoMigrate(dst...)
}

func DB() *gorm.DB {
	return db
}

func Initialize() (err error) {
	obj := &Object{
		Driver: config.GetDbDriver(),
		DSN:    config.GetDSN(),
	}

	db, err = obj.open()
	return
}
