package config

import "strconv"

type ServerConfig struct {
	NetworkInterface string      `toml:"network_interface"`
	BindInterface    bool        `toml:"bind_interface"`
	UseInterfaceIp   bool        `toml:"use_interface_ip"`
	Http             *HttpConfig `toml:"http"`
}

type HttpConfig struct {
	Host    string `toml:"host"`
	Port    int    `toml:"port"`
	Charset string `toml:"charset"`
	Gzip    bool   `toml:"gzip"`
	PProf   bool   `toml:"pprof"`
}

func GetHttp() *HttpConfig {
	return config.Server.Http
}

func (c *HttpConfig) GetHost() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}
