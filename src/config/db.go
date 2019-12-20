package config

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris/v12"
)

type DbConfig struct {
	Drive                string `toml:"drive"`
	Host                 string `toml:"host"`
	Port                 string `toml:"port"`
	User                 string `toml:"user"`
	Password             string `toml:"password"`
	Charset              string `toml:"charset"`
	Database             string `toml:"database"`
	Timeout              int    `toml:"timeout" json:"timeout"`
	MaxOpenConns         int    `toml:"max_open_conns" json:"max_open_conns"`
	MaxIdleConns         int    `toml:"max_idle_conns" json:"max_idle_conns"`
	MaxConnTtlMaxConnTtl int    `toml:"max_conn_ttl" json:"max_conn_ttl"`
	Debug                bool   `toml:"debug"`
}

func InitDB(c *DbConfig) (err error) {
	if Log == nil || Log.Logger == nil {
		return errors.New("logger uninitialized")
	}
	DB, err = gorm.Open(c.Drive, c.User+":"+c.Password+"@/"+c.Database+"?charset=utf8&parseTime=True&loc=Local") //protocol("+c.Host+":"+c.Port+")
	if err != nil {
		Log.LogInfo("orm failed to initialized: ", err)
	}
	iris.RegisterOnInterrupt(func() {
		_ = DB.Close()
	})
	return err
}

func GetDB() *DbConfig {
	return config.DB
}
