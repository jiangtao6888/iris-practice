package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func (r *router) RegHttpHandler(app *iris.Application) {
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML(" <h1>hi, the server is running</h1>")
	})
}

func (r *router) GetIdentifier(ctx context.Context) string {
	return "role"
}

var Router *router

type router struct{}
