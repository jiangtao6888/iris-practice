package config

import (
	"github.com/BurntSushi/toml"
)

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
