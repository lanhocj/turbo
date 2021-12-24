package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net"
)

type Object struct {
	Database struct {
		Driver   string `yaml:"driver" json:"driver,omitempty"`
		URI      string `yaml:"uri" json:"uri,omitempty"`
		Host     string `yaml:"host" json:"host,omitempty"`
		Port     string `yaml:"port" json:"port,omitempty"`
		User     string `yaml:"user" json:"user"`
		Password string `yaml:"password" json:"password"`
		Table    string `yaml:"table" json:"table"`
	} `yaml:"database" json:"database"`

	Http struct {
		Host string `yaml:"host" json:"host"`
		Port string `yaml:"port" json:"port"`
	} `yaml:"http" json:"http"`
}

var Config = &Object{}

func New() *Object { return Config }

func (c *Object) LoadFile(in string) error {
	tmp, err := ioutil.ReadFile(in)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(tmp, c); err != nil {
		return err
	}

	return nil
}

func (c *Object) GetAddr() string     { return net.JoinHostPort(c.Http.Host, c.Http.Port) }
func (c *Object) GetDbDriver() string { return Config.Database.Driver }

func (c *Object) GetDSN() string {
	if len(c.Database.URI) <= 0 {
		// user:password@tcp(localhost:3306)/test
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
			c.Database.User,
			c.Database.Password,
			c.Database.Host,
			c.Database.Port,
			c.Database.Table)
	}

	return c.Database.URI
}

// GetAddr returns addr..
func GetAddr() string     { return Config.GetAddr() }
func GetDbDriver() string { return Config.GetDbDriver() }
func GetDSN() string      { return Config.GetDSN() }
