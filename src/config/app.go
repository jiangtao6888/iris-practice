package config

import (
	"database/sql"
	"github.com/go-redis/redis/v7"
)

var config *Config
var HttpServer *Server
var AccLog *AccessLog
var DB *sql.DB
var Log *Logger
var Cache *redis.Client
