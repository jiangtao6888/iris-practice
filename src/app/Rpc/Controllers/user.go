package Controllers

import (
	"context"
	"iris/app/Code"
	"iris/app/Helpers"
	"iris/app/Service"
	"iris/config"
	"iris/libraries/proto"
)

var User = &user{}

type user struct{}

func (c *user) GetUserInfo(ctx context.Context, req *proto.UserRequest) (rsp *proto.UserResponse, err error) {
	config.Log.LogInfo(req)
	rsp = &proto.UserResponse{
		Code:    Code.SuccessCode,
		Message: Code.Message[Code.SuccessCode],
		Data:    &proto.User{},
	}
	user, err := Service.User.GetUserInfo(req.UserId)
	if err != nil {
		rsp.Code = Code.ErrorCode
		return
	}
	if err = Helpers.ConvertStruct(user, rsp.Data); err != nil {
		rsp.Code = Code.ErrorCode
		return
	}
	return
}
