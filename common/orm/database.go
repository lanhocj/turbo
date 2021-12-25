package orm

import (
	"github.com/laamho/turbo/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Object struct {
	db          *gorm.DB
	Driver, DSN string
}

var DB *gorm.DB

func openMysql(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (d *Object) Open() *gorm.DB {
	switch d.Driver {
	case "mysql":
		return openMysql(d.DSN)
	default:
		return nil
	}
}

func Init() {
	o := &Object{
		Driver: config.GetDbDriver(),
		DSN:    config.GetDSN(),
	}

	log.Printf("Driver: %s, Dsn: %s", config.GetDbDriver(), config.GetDSN())

	DB = o.Open()

	if err := DB.AutoMigrate(&User{}, &Node{}); err != nil {
		log.Fatal(err)
	}
}
