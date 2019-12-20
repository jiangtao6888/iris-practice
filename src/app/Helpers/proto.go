package Helpers

import (
	"errors"
	"github.com/gogo/protobuf/proto"
	"github.com/kataras/iris/v12/context"
)

var Proto= &protoCoder{}

type protoCoder struct{}

func (c *protoCoder) Unmarshal(data []byte, v interface{}) error {
	pb, ok := v.(proto.Message)

	if !ok {
		return errors.New("invalid protobuf message")
	}

	return proto.Unmarshal(data, pb)
}

func (c *protoCoder) Marshal(v interface{}) ([]byte, error) {
	pb, ok := v.(proto.Message)

	if !ok {
		return nil, errors.New("invalid protobuf message")
	}

	return proto.Marshal(pb)
}

func (c *protoCoder) DecodeIrisReq(ctx context.Context, v interface{}) error {
	return ctx.UnmarshalBody(v, c)
}

func (c *protoCoder) SendIrisReply(ctx context.Context, v interface{}) error {
	ctx.ContentType(context.ContentBinaryHeaderValue)
	ctx.Header(EncodingHeader, EncodingProtobuf)

	data, err := c.Marshal(v)

	if err != nil {
		return err
	}

	_, err = ctx.Write(data)
	return err
}
