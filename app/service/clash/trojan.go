package clash

import (
	"github.com/laamho/turbo/common"
	"gopkg.in/yaml.v2"
	"strconv"
)

type ProxyConfig struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Server   string `yaml:"server"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	Udp      bool   `yaml:"udp"`
}

type ProxyGroupConfig struct {
	Name    string   `yaml:"name"`
	Type    string   `yaml:"type"`
	Proxies []string `yaml:"proxies"`
}

type RuleProvider struct {
	Type     string `yaml:"type"`
	Behavior string `yaml:"behavior"`
	Url      string `yaml:"url"`
	Path     string `yaml:"path"`
	Interval int    `yaml:"interval"`
}

type Config struct {
	Port        uint               `yaml:"port"`
	SocksPort   uint               `yaml:"socks-port"`
	Mode        string             `yaml:"mode"`
	LogLevel    string             `yaml:"log-level"`
	Proxies     []ProxyConfig      `yaml:"proxies"`
	ProxyGroups []ProxyGroupConfig `yaml:"proxy-groups"`

	RuleProviders map[string]RuleProvider `yaml:"rule-providers"`
	Rules         []string                `yaml:"rules"`
}

func NewProxyGroup(n, t string) ProxyGroupConfig {
	return ProxyGroupConfig{
		Name: n,
		Type: t,
	}
}

func (c *Config) AddProxy(t, n, addr, port, password string, udp bool) {
	parsedPort, _ := strconv.Atoi(port)

	p := ProxyConfig{
		Name:     n,
		Type:     t,
		Server:   addr,
		Port:     parsedPort,
		Password: password,
		Udp:      udp,
	}

	c.Proxies = append(c.Proxies, p)
}

func (c *Config) String() string {

	reject := NewProxyGroup("FINAL", "select")
	reject.Proxies = []string{"REJECT"}

	proxyGroup := NewProxyGroup("PROXY", "select")

	for _, proxy := range c.Proxies {
		proxyGroup.Proxies = append(proxyGroup.Proxies, proxy.Name)
	}

	c.ProxyGroups = append(c.ProxyGroups, proxyGroup)
	c.ProxyGroups = append(c.ProxyGroups, reject)

	c.RuleProviders = map[string]RuleProvider{
		"reject":       {Type: "http", Behavior: "domain", Url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/reject.txt", Path: "./ruleset/reject.yaml", Interval: 86400},
		"direct":       {Type: "http", Behavior: "domain", Url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/direct.txt", Path: "./ruleset/direct.yaml", Interval: 86400},
		"lancidr":      {Type: "http", Behavior: "domain", Url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/lancidr.txt", Path: "./ruleset/lancidr.yaml", Interval: 86400},
		"cncidr":       {Type: "http", Behavior: "domain", Url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/cncidr.txt", Path: "./ruleset/cncidr.yaml", Interval: 86400},
		"telegramcidr": {Type: "http", Behavior: "domain", Url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/telegramcidr.txt", Path: "./ruleset/telegramcidr.yaml", Interval: 86400},
		"google":       {Type: "http", Behavior: "domain", Url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/google.txt", Path: "./ruleset/google.yaml", Interval: 86400},
	}

	c.Rules = []string{
		"RULE-SET,reject,REJECT",
		"RULE-SET,direct,DIRECT",
		"RULE-SET,lancidr,DIRECT",
		"RULE-SET,cncidr,DIRECT",
		"RULE-SET,google,PROXY",
		"RULE-SET,telegramcidr,PROXY",
		"GEOIP,CN,DIRECT,no-resolve",
		"MATCH,PROXY",
	}

	out, err := yaml.Marshal(c)
	if err != nil {
		common.Silent(err)
	}

	return string(out)
}

func Default() *Config {
	c := new(Config)
	c.Port = uint(7890)
	c.SocksPort = uint(7891)
	c.LogLevel = "info"
	c.Mode = "Rule"
	return c
}
