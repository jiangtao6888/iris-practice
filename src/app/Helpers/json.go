package Helpers

import (
	"bytes"
	"encoding/json"
	"github.com/kataras/iris/v12/context"
	"net/http"
)

var Json = &jsonCoder{}

type jsonCoder struct{}

func (c *jsonCoder) Unmarshal(data []byte, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	return decoder.Decode(v)
}

func (c *jsonCoder) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (c *jsonCoder) UnescapeMarshal(v interface{}) ([]byte, error) {
	bs, err := c.Marshal(v)

	if err == nil {
		bs = bytes.Replace(bs, []byte(`\\n`), []byte(`\n`), -1)
		bs = bytes.Replace(bs, []byte(`\\'`), []byte(`\'`), -1)
		bs = bytes.Replace(bs, []byte(`\\"`), []byte(`\"`), -1)
	}

	return bs, err
}

func (c *jsonCoder) DecodeIrisReq(ctx context.Context, v interface{}) error {
	return ctx.UnmarshalBody(v, c)
}

func (c *jsonCoder) SendIrisReply(ctx context.Context, v interface{}) error {
	option := context.JSON{}

	ctx.ContentType(context.ContentJSONHeaderValue)
	ctx.Header(EncodingHeader, EncodingJson)

	if s := ctx.Values().Get(CtxRspUnescapeKey); s != nil && s.(bool) {
		body, err := c.UnescapeMarshal(v)

		if err != nil {
			ctx.StatusCode(http.StatusInternalServerError)
			return err
		}

		_, err = ctx.Write(body)

		if err != nil {
			ctx.StatusCode(http.StatusInternalServerError)
			return err
		}

		return nil
	} else {
		_, err := ctx.JSON(v, option)
		return err
	}
}
