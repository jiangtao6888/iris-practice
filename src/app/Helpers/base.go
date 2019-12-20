package Helpers

import "github.com/kataras/iris/v12/context"

const (
	EncodingHeader   = "Protocol-Encoding"
	EncodingProtobuf = "protobuf"
	EncodingJson     = "json"
)

type ICoder interface {
	Unmarshal(data []byte, v interface{}) error
	Marshal(v interface{}) ([]byte, error)
	DecodeIrisReq(ctx context.Context, v interface{}) error
	SendIrisReply(ctx context.Context, v interface{}) error
}
