package Controllers

import (
	"context"
	"iris/app/Service"
	"iris/config"
	"iris/libraries/proto"
)

var User = &user{}

type user struct{}

func (c *user) GetUserInfo(ctx context.Context, req *proto.UserRequest) (rsp *proto.UserResponse, err error) {
	config.Log.LogInfo(req)
	rsp = &proto.UserResponse{}
	rsp, err = Service.User.GetUserInfo(req.UserId)
	return
}
