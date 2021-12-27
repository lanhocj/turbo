package config

import (
	"encoding/base64"
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

	App struct {
		Title       string `yaml:"title"`
		Secret      string `yaml:"secret"`
		Description string `yaml:"description"`
		Copyright   string `yaml:"copyright"`
	} `yaml:"app"`
}

var cnf = &Object{}

func New() *Object { return cnf }

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
func (c *Object) GetDbDriver() string { return cnf.Database.Driver }
func (c *Object) getSecret() string   { return cnf.App.Secret }

func (c *Object) GetDSN() string {
	if len(c.Database.URI) <= 0 {
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
			c.Database.User,
			c.Database.Password,
			c.Database.Host,
			c.Database.Port,
			c.Database.Table)
	}

	return c.Database.URI
}

// GetAddr returns addr..
func GetAddr() string     { return cnf.GetAddr() }
func GetDbDriver() string { return cnf.GetDbDriver() }
func GetDSN() string      { return cnf.GetDSN() }
func Copy() *Object       { return cnf }

func GetSecret() []byte {
	dec, _ := base64.StdEncoding.DecodeString(cnf.getSecret())
	return dec
}
