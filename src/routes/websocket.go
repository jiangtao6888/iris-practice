package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
)

func (r *router) WebsocketHandler(server *neffos.Server, app *iris.Application) {
	app.Any("/echo", websocket.Handler(server))
}
