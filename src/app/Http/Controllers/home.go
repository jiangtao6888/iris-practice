package Controllers

import (
	"github.com/kataras/iris/v12/context"
	"iris/app/Http/Auth"
)

var Home = &home{}

type home struct{}

func (l *home) Index(ctx context.Context) {
	jwt := Auth.TokenHandler(ctx)
	ctx.HTML(" <h1>hi, the server is running</h1><br>Jwt:" + jwt)
	return
}
