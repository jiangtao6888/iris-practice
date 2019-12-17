package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Server *ServerConfig `toml:"server"`
	DB     *DbConfig     `toml:"db"`
	Cache  *CacheConfig  `toml:"cache"`
	SysLog *SysLogConfig `toml:"log"`
}



func InitConfig(file string) error {
	config = &Config{
		Server: &ServerConfig{
			Http: &HttpConfig{},
		},
	}
	if _, err := toml.DecodeFile(file, config); err != nil {
		return err
	}
	return nil
}
