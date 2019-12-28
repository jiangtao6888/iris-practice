package config

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"strconv"
	"strings"
	"time"
)

func (c *CacheConfig) CacheConfig() *redis.Options {
	config := &redis.Options{
		Addr:     c.Host + ":" + strconv.Itoa(c.Port),
		Password: c.Password,
		PoolSize: c.PoolSize,
	}
	return config
}

func InitCache(c *CacheConfig) error {
	Client = redis.NewClient(c.CacheConfig())
	return nil
}

func GetCache() *CacheConfig {
	return config.Cache
}

type RedisCache struct {
}

func (c *RedisCache) GetKey(prefix string, items ...interface{}) string {
	format := prefix + strings.Repeat(":%v", len(items))
	return fmt.Sprintf(format, items...)
}

func (c *RedisCache) SetArray(key string, data interface{}, expire time.Duration) error {
	b, err := json.Marshal(data)
	if err != nil {
		Log.LogInfo(" Marshal redis error", err)
		return err
	}
	return Client.Set(key, string(b), expire).Err()
}

func (c *RedisCache) GetArray(key string) string {
	val, err := Client.Get(key).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		return ""
	} else {
		return val
	}
}

func (c *RedisCache) Expire(key string, time time.Duration) error {
	return Client.Expire(key, time).Err()
}
