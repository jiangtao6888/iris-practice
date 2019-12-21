package Controllers

import (
	"github.com/kataras/iris/v12/context"
	"iris/app/Code"
	"iris/app/Helpers"
	"iris/app/Service"
	"iris/config"
)

var Login = &login{}

type login struct{}

func (l *login) Login(ctx context.Context) {

	rsp, err := Service.Login.Login(ctx)

	if err == false {
		config.Log.LogInfo("用户信息不正确", rsp)
		Helpers.Error(ctx, Code.UserNotExist)
		return
	}

	Helpers.SendRsp(ctx, rsp)
	return
}
