package Service

import (
	"iris/app/Code"
	"iris/app/Exceptions"
	"iris/app/Helpers"
	"iris/app/Models"
	"iris/libraries/proto"
)

var User = &user{}

type user struct{}

func (s *user) GetUserInfo(UserId int64) (*proto.UserResponse, error) {

	user, err := Models.GetUserInfo(UserId)

	if err != nil {
		return nil, Exceptions.New(Code.ErrorCode, "用户不存在")
	}

	rsp := &proto.UserResponse{
		Code:    Code.SuccessCode,
		Message: Code.Message[Code.SuccessCode],
		Data:    &proto.User{},
	}

	if err := Helpers.ConvertStruct(user, rsp.Data); err != nil {
		return nil, Exceptions.New(Code.ErrorCode, "转化失败")
	}

	return rsp, nil
}
