package config

import (
	stdContext "context"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/pprof"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type AccessLog struct {
	Logger *logrus.Logger
}

func GetAccLog() *HttpLogConfig {
	return config.Server.Http.Log
}

func GetHttp() *ServerConfig {
	return config.Server
}

func InitHttpServer(c *ServerConfig, route Router) error {
	HttpServer = NewHttpServer(c, AccLog, route)
	HttpServer.Start(c)
	return nil
}

func GetClientIp(ctx context.Context) string {
	xForwarded := ctx.GetHeader("X-Forwarded-For")

	if ip := strings.TrimSpace(strings.Split(xForwarded, ",")[0]); ip != "" {
		return ip
	}

	if xReal := strings.TrimSpace(ctx.GetHeader("X-Real-Ip")); xReal != "" {
		return xReal
	}

	return ctx.RemoteAddr()
}

func NewHttpServer(c *ServerConfig, OutFile *AccessLog, router Router) *Server {
	server := &Server{config: c, router: router}
	server.app = iris.New()
	server.ctx, server.canceler = stdContext.WithCancel(stdContext.Background())

	server.app.Use(server.Recovery)
	server.app.Use(server.AccessLog)

	// enable gzip
	if c.Http.Gzip {
		server.app.Use(iris.Gzip)
	}

	// enable pprof
	if c.Http.PProf {
		server.app.Any("/debug/pprof", pprof.New())
		server.app.Any("/debug/pprof/{action:path}", pprof.New())
	}

	// set logger
	server.app.Logger().SetLevel(string(c.Http.Log.Level))
	server.app.Logger().SetTimeFormat(c.Http.Log.TimeFormat)
	server.app.Logger().SetOutput(OutFile)
	server.app.Logger().Printer.IsTerminal = c.Http.Log.Color

	// set route
	server.router.RegHttpHandler(server.app)

	return server
}

func (l *HttpLogConfig) directory() string {
	return l.Directory
}

func (l *HttpLogConfig) level() uint32 {
	return l.Level
}

func (l *HttpLogConfig) new(f *os.File) {
	AccLog = &AccessLog{Logger: logrus.New()}
	AccLog.Logger.Out = f
}

func (l *HttpLogConfig) name() string {
	return "access-" + time.Now().Format("20060102") + ".log"
}

func (l *AccessLog) LogInfo(msg ...interface{}) {
	l.Logger.Info(msg)
}

func (l *AccessLog) LogInfoFields(fields Fields, msg interface{}) {
	l.Logger.WithFields(logrus.Fields(fields)).Info(msg)
}

func (l *AccessLog) Write(p []byte) (n int, err error) {
	return
}

func (c *ServerConfig) GetHost() string {
	return c.Http.Host + ":" + strconv.Itoa(c.Http.Port)
}

func (s *Server) Running() bool {
	select {
	case <-s.ctx.Done():
		return false
	default:
		return true
	}
}

func (s *Server) Start(c *ServerConfig) {
	go func() {
		var runner iris.Runner

		_ = s.app.Run(iris.Addr(c.GetHost()), iris.WithoutServerError(iris.ErrServerClosed))

		err := s.app.Run(runner, iris.WithConfiguration(iris.Configuration{
			DisableStartupLog:                 true,
			DisableInterruptHandler:           true,
			DisableBodyConsumptionOnUnmarshal: true,
			Charset:                           s.config.Charset,
		}))

		if err != nil && s.Running() {
			Log.LogInfo("can't serve at <%s> | error: %s", c.GetHost(), err)
		}
	}()

}

// recovery panic (500)
func (s *Server) Recovery(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			if ctx.IsStopped() {
				return
			}

			var stacktrace string

			for i := 1; ; i++ {
				_, f, l, got := runtime.Caller(i)

				if !got {
					break
				}

				stacktrace += fmt.Sprintf("%s:%d\n", f, l)
			}

			request := fmt.Sprintf("%v %s %s %s", strconv.Itoa(ctx.GetStatusCode()), GetClientIp(ctx), ctx.Method(), ctx.Path())
			Log.LogInfo(fmt.Sprintf("recovered panic:\nRequest: %s\nTrace: %s\n%s", request, err, stacktrace))
			ctx.StatusCode(500)
			ctx.StopExecution()
		}
	}()

	ctx.Next()
}

// record access log
func (s *Server) AccessLog(ctx context.Context) {
	start := time.Now()
	ctx.Next()
	idf := s.router.GetIdentifier(ctx)
	statusCode, useTime, clientIp := ctx.GetStatusCode(), time.Since(start), GetClientIp(ctx)
	uri, method, userAgent := ctx.Request().URL.RequestURI(), ctx.Method(), ctx.GetHeader("User-Agent")
	filed := Fields{
		"statusCode": statusCode,
		"clientIp":   clientIp,
		"useTime":    useTime,
		"method":     method,
		"uri":        uri,
		"userAgent":  userAgent,
		"idf":        idf,
	}
	AccLog.LogInfoFields(filed, "request")
}
