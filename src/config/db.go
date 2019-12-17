package config

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"time"
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

func InitDB(c *DbConfig) error {
	if Log == nil || Log.Logger == nil {
		return errors.New("logger uninitialized")
	}
	DB, _ = sql.Open(c.Drive, c.User+":"+c.Password+"@tcp("+c.Host+":"+c.Port+")"+c.Database)
	DB.SetMaxOpenConns(c.MaxOpenConns)
	DB.SetMaxIdleConns(c.MaxIdleConns)
	DB.SetConnMaxLifetime(time.Duration(c.MaxConnTtlMaxConnTtl))

	return nil
}

func GetDB() *DbConfig {
	return config.DB
}
