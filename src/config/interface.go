package config

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"os"
)

type logInterface interface {
	directory() string
	level() uint32
	name() string
	new(f *os.File)
}

type Router interface {
	RegHttpHandler(app *iris.Application)
	GetIdentifier(ctx context.Context) string
}

type RpcRouter interface {
	RegRpcService(server *RpcServer)
}
