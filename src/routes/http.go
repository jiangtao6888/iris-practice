package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"iris/app/Http/Controllers"
	middleware "iris/app/Http/Middleware"
)

func (r *router) RegHttpHandler(app *iris.Application) {
	app.Use(middleware.JsonCoder)

	app.Get("/", Controllers.Home.Index)

	user := app.Party("user")
	{
		user.Post("/infos", Controllers.User.GetUserInfo)
	}
}

func (r *router) GetIdentifier(ctx context.Context) string {
	return "role"
}

var Router *router

type router struct{}
