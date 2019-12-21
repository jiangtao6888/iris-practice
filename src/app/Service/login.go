package Service

import (
	"github.com/kataras/iris/v12/context"
	"iris/app/Code"
	"iris/app/Helpers"
	"iris/app/Http/Auth"
	"iris/app/Models"
	"iris/libraries/proto"
)

var Login = &login{}

type login struct{}

func (s *login) Login(ctx context.Context) (rsp *proto.LoginResponse, err bool) {

	req := &proto.LoginRequest{}

	if !Helpers.DecodeReq(ctx, req) {
		return nil, false
	}
	u, e := Models.GetUserInfoByUserName(req.Username)

	if e != nil {
		return nil, false
	}

	info := &proto.Login{
		Id:       u.Id,
		Username: u.UserName,
		Token:    Auth.GetToken(u.Id),
	}

	rsp = &proto.LoginResponse{
		Code:    Code.SuccessCode,
		Message: Code.Message[Code.SuccessCode],
		Data:    info,
	}

	return rsp, true
}
