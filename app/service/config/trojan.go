package config

type ProxyConfig struct {
}

type ProxyGroupConfig struct {
}

type ClashClientConfig struct {
	Port        int8               `yaml:"port"`
	SocksPort   int8               `yaml:"socks-port"`
	Mode        string             `yaml:"mode"`
	LogLevel    string             `yaml:"log-level"`
	Proxies     []ProxyConfig      `yaml:"proxies"`
	ProxyGroups []ProxyGroupConfig `yaml:"proxy-groups"`
}
