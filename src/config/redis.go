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

func (c *CacheConfig) CacheConfig() *redis.Options {
	config := &redis.Options{
		Addr:     c.Host + ":" + strconv.Itoa(c.Port),
		Password: c.Password,
		PoolSize: c.PoolSize,
	}
	return config
}

func InitCache(c *CacheConfig) error {
	Cache = redis.NewClient(c.CacheConfig())
	return nil
}

func GetCache() *CacheConfig {
	return config.Cache
}
