package Controllers

import (
	"github.com/kataras/iris/v12/context"
	"iris/app/Code"
	"iris/app/Helpers"
	"iris/app/Service"
	"iris/config"
	"iris/libraries/proto"
)

var User = &user{}

type user struct{}

func (u *user) UserInfo(ctx context.Context) {

	req := &proto.UserRequest{}

	if !Helpers.DecodeReq(ctx, req) {
		return
	}

	user, err := Service.User.GetUserInfo(req.UserId)
	if err != nil {
		config.Log.LogInfo("用户信息为空")
		Helpers.Error(ctx, Code.ErrorCode, Code.Message[Code.ErrorCode])
		return
	}

	rsp := &proto.UserResponse{
		Code:    Code.SuccessCode,
		Message: Code.Message[Code.SuccessCode],
		Data:    &proto.User{},
	}

	if err := Helpers.ConvertStruct(user, rsp.Data); err != nil {
		Helpers.Error(ctx, Code.ErrorCode, Code.Message[Code.ErrorCode])
		return
	}
	Helpers.SendRsp(ctx, rsp)
	return
}
