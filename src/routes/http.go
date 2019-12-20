package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"iris/app/Http/Auth"
	"iris/app/Http/Controllers"
	middleware "iris/app/Http/Middleware"
)

func (r *router) RegHttpHandler(app *iris.Application) {
	app.Get("/", func(ctx iris.Context) {
		jwt := Auth.TokenHandler(ctx)
		ctx.HTML(" <h1>hi, the server is running</h1><br>Jwt:" + jwt)
	})
	user := app.Party("user", middleware.JsonCoder)
	{
		user.Use(Auth.AuthenticatedHandler)
		user.Post("/info", Controllers.User.UserInfo)
	}
}

func (r *router) GetIdentifier(ctx context.Context) string {
	return "role"
}

var Router *router

type router struct{}
