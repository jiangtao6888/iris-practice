package Helpers

import (
	"encoding/json"
	"github.com/kataras/iris/v12/context"
	"iris/app/Code"
	"iris/config"
	"iris/libraries/proto"
)

func ConvertStruct(a interface{}, b interface{}) error {
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, b)
	if err != nil {
		return err
	}
	return nil
}

func Error(ctx context.Context, errCode int64) {
	//data := make(map[int32]string)
	SendRsp(ctx, &proto.Response{Code: errCode, Message: Code.Message[errCode]})
}

func SendRsp(ctx context.Context, rsp interface{}) {
	err := GetCtxCoder(ctx).SendIrisReply(ctx, rsp)
	if err != nil {
		config.Log.LogInfo("can't send http response | error: ", err)
	}
}
