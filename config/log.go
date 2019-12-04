package config

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type LogConfig struct {
	Level uint32 `toml:"level"`
}

func InitLogger(config *LogConfig) (err error) {
	// 设置日志格式为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	file := time.Now().Format("20060102") + ".log" //文件名
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		panic(err)
	}
	logrus.SetOutput(logFile)

	// 设置日志级别为warn以上
	logrus.SetLevel(logrus.Level(config.Level))
	return err
}

func GetLog() *LogConfig {
	return config.Log
}

func Log() *logrus.Logger {
	return logrus.New()
}

func LogInfo(fields logrus.Fields, msg interface{}) {
	Log().WithFields(fields).Info(msg)
}
