package Auth

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"iris/app/Code"
	"iris/app/Helpers"
	"iris/config"
)

func AuthenticatedHandler(ctx iris.Context) {
	Server:= jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetJwt().Secret), nil
		},
		Extractor:     jwt.FromAuthHeader,
		SigningMethod: jwt.SigningMethodHS256,
	})

	if err := Server.CheckJWT(ctx); err != nil {
		config.Log.LogInfo("user.Valid:false")
		Helpers.Error(ctx,Code.AuthenticatedCode)
		return
	}
	user := ctx.Values().Get("jwt").(*jwt.Token)
	config.Log.LogInfo("user.Valid:",user.Valid)
	ctx.Next()
}