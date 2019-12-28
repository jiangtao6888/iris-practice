package config

import (
	stdContext "context"
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sync"
)

var config *Config
var HttpServer *Server
var AccLog *AccessLog
var DB *gorm.DB
var Log *Logger
var Client *redis.Client
var Rpc *RpcServer
var Cache *RedisCache
var KafkaProducer *Producer

type Config struct {
	Server     *ServerConfig   `toml:"server"`
	DB         *DbConfig       `toml:"db"`
	Cache      *CacheConfig    `toml:"cache"`
	SysLog     *SysLogConfig   `toml:"log"`
	Jwt        *JwtConfig      `toml:"jwt"`
	RpcConfig  *RpcConfig      `toml:"rpc"`
	KaProducer *ProducerConfig `toml:"kafka"`
}

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

type ServerConfig struct {
	NetworkInterface string      `toml:"network_interface"`
	BindInterface    bool        `toml:"bind_interface"`
	UseInterfaceIp   bool        `toml:"use_interface_ip"`
	Http             *HttpConfig `toml:"http"`
	Charset          string      `toml:"charset"`
}

type HttpConfig struct {
	Host    string         `toml:"host"`
	Port    int            `toml:"port"`
	Charset string         `toml:"charset"`
	Gzip    bool           `toml:"gzip"`
	PProf   bool           `toml:"pprof"`
	Log     *HttpLogConfig `toml:"log"`
}

type HttpLogConfig struct {
	Level      uint32 `toml:"level"`
	Directory  string `toml:"directory"`
	TimeFormat string `toml:"time_format"`
	Color      bool   `toml:"color"`
}

type Server struct {
	sync.Mutex
	config   *ServerConfig
	router   Router
	app      *iris.Application
	ctx      stdContext.Context
	canceler func()
}

type SysLogConfig struct {
	Level     uint32 `toml:"level"`
	Directory string `toml:"directory"`
}

type CacheConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	Database int    `toml:"database"`
	PoolSize int    `toml:"poolsize"`
}

type JwtConfig struct {
	Secret string `toml:"secret"`
}

type RpcConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type RpcServer struct {
	router   RpcRouter
	Server   *grpc.Server
	ctx      stdContext.Context
	Log      *AccessLog
	canceler func()
}

type ProducerConfig struct {
	Brokers []string `toml:"brokers"`
}

type Producer struct {
	c      *ProducerConfig
	client sarama.AsyncProducer
	log    *Logger
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup
}
