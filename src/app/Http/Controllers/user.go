package Controllers

import (
	"github.com/kataras/iris/v12/context"
	"iris/app/Code"
	"iris/app/Helpers"
	"iris/app/Service"
	"iris/libraries/proto"
)

var User = &user{}

type user struct{}

func (u *user) UserInfo(ctx context.Context) {

	req := &proto.UserRequest{}

	if !Helpers.DecodeReq(ctx, req) {
		Helpers.Error(ctx, Code.ErrorCode)
		return
	}

	rsp, err := Service.User.GetUserInfo(req.UserId)
	if err != nil {
		Helpers.Error(ctx, Code.ErrorCode)
		return
	}
	Helpers.SendRsp(ctx, rsp)
	return
}
