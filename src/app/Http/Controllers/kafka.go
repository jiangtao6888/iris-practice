package Controllers

import (
	"github.com/kataras/iris/v12/context"
	"iris/app/Code"
	"iris/app/Helpers"
	"iris/app/Service"
	"iris/libraries/proto"
)

var Kafka = &kafka{}

type kafka struct{}

func (k *kafka) Send(ctx context.Context) {
	topic := "iris.test"
	err := Service.Kafka.Send(topic, `{"user_id":1}`)
	if err != nil {
		Helpers.Error(ctx, Code.ErrorCode)
		return
	}
	rsp := proto.Response{Code: Code.SuccessCode, Message: "Success !"}
	Helpers.SendRsp(ctx, rsp)
	return
}
