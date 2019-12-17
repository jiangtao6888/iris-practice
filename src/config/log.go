package config

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)



type Logger struct {
	Logger *logrus.Logger
}

func (l *Logger) Write(p []byte) (n int, err error) {
	return
}

type logInterface interface {
	directory() string
	level() uint32
	name() string
	new(f *os.File)
}

func (l *SysLogConfig) name() string {
	return "system-" + time.Now().Format("20060102") + ".log"
}

func (l *SysLogConfig) directory() string {
	return l.Directory
}

func (l *SysLogConfig) level() uint32 {
	return l.Level
}

func (l *SysLogConfig) new(f *os.File) {
	Log = &Logger{Logger: logrus.New()}
	Log.Logger.Out = f
}

type SysLogConfig struct {
	Level     uint32 `toml:"level"`
	Directory string `toml:"directory"`
}

func InitLogger(l logInterface) (err error) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	file := l.directory() + l.name() //文件名
	logFile, e := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if e != nil {
		return e
	}
	// 设置日志级别为warn以上
	logrus.SetLevel(logrus.Level(l.level()))
	l.new(logFile)
	return err
}

func GetLog() *SysLogConfig {
	return config.SysLog
}

type Fields map[string]interface{}

func (l *Logger) LogInfo(msg ...interface{}) {
	Log.Logger.Info(msg)
}

func (l *Logger) LogInfoFields(fields Fields, msg interface{}) {
	Log.Logger.WithFields(logrus.Fields(fields)).Info(msg)
}
