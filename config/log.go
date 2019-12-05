package config

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Logger *logrus.Logger

type LogConfig struct {
	Level     uint32 `toml:"level"`
	Directory string `toml:"directory"`
}

func InitLogger(config *LogConfig) (err error) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	file := config.Directory + time.Now().Format("20060102") + ".log" //文件名
	logFile, e := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if e != nil {
		return e
	}
	// 设置日志级别为warn以上
	logrus.SetLevel(logrus.Level(config.Level))
	Logger = LoggerNew()
	Logger.Out = logFile
	return err
}

func GetLog() *LogConfig {
	return config.Log
}

func LoggerNew() *logrus.Logger {
	return logrus.New()
}

type Fields map[string]interface{}

func LogInfo(msg ...interface{}) {
	Logger.Info(msg)
}

func LogInfoFields(fields Fields, msg interface{}) {
	Logger.WithFields(logrus.Fields(fields)).Info(msg)
}
