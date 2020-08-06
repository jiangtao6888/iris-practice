package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/kataras/iris/v12/context"
	"google.golang.org/protobuf/encoding/protojson"
	"iris/libraries/proto"
)

var User = &user{}

type user struct{}

func (u *user) GetUserInfo(ctx context.Context) {
	rsp := &proto.Response{
		Code:    200,
		Message: "success",
	}

	list := &proto.List{}

	for i := 0; i < 10; i++ {
		data := &proto.User{
			Id:       int32(i),
			Name:     "knight",
			Username: "knight6888",
			RoleId:   1,
		}
		list.User = append(list.User, data)
	}

	ret, err := ptypes.MarshalAny(list)

	if err != nil {
		fmt.Println(err)
		return
	}

	rsp.Data = ret
	p := &protojson.MarshalOptions{
		AllowPartial:  true,
		UseProtoNames: true,
	}
	bytes, _ := p.Marshal(rsp)

	var r map[string]interface{}
	_ = json.Unmarshal(bytes, &r)

	_, _ = ctx.JSON(r)
	return
}
