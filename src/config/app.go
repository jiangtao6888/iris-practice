package config

import (
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
)

var config *Config
var HttpServer *Server
var AccLog *AccessLog
var DB *gorm.DB
var Log *Logger
var Cache *redis.Client
