package middleware

import (
	"github.com/kataras/iris/v12/context"
	"iris/app/Helpers"
)

func JsonCoder(ctx context.Context) {
	Helpers.SetCtxCoder(ctx, Helpers.EncodingJson)
	ctx.Next()
}

func ProtoCoder(ctx context.Context) {
	Helpers.SetCtxCoder(ctx, Helpers.EncodingProtobuf)
	ctx.Next()
}