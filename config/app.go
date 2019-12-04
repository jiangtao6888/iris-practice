package config

import (
	"github.com/kataras/iris/v12"
)
func NewApp() {
	app := iris.New()


	app.Get("/", func(ctx iris.Context) {
		ctx.HTML(" <h1>hi, the server is running</h1>")
	})
	_ = app.Run(iris.Addr(GetHttp().GetHost()), iris.WithoutServerError(iris.ErrServerClosed))
}
