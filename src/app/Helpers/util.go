package Helpers

import (
	"github.com/kataras/iris/v12/context"
	"iris/config"
)

func DecodeReq(ctx context.Context, req interface{}) bool {
	if err := GetCtxCoder(ctx).DecodeIrisReq(ctx, req); err != nil {
		config.Log.LogInfo("invalid parameter | req:  | error: ", req, err)
		return false
	}
	return true
}

func GetCtxCoder(ctx context.Context) ICoder {
	name := ctx.Values().GetString(CtxCoderKey)
	if name == EncodingProtobuf {
		return Proto
	}
	return Json
}

func SetCtxCoder(ctx context.Context, encoding string) {
	if encoding == EncodingProtobuf || encoding == EncodingJson {
		ctx.Values().Set(CtxCoderKey, encoding)
	}
}
