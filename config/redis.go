package config

import (
	"github.com/go-redis/redis/v7"
	"strconv"
	_ "sync"
)

type CacheConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	Database int    `toml:"database"`
	PoolSize int    `toml:"poolsize"`
}

func Redis() *redis.Client {
	return redis.NewClient(config.GetConfig())
}

//var Cache = redis.NewClient(cache.GetConfig())

func (c *Config) GetConfig() *redis.Options {
	config := &redis.Options{
		Addr:     c.Cache.Host + ":" + strconv.Itoa(c.Cache.Port),
		Password: c.Cache.Password,
		PoolSize: c.Cache.PoolSize,
	}
	return config
}


