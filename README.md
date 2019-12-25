# iris-practice
iris练习使用


# 使用

```
make p
make dep
make iris
```

这是grpc调用的测试client
```
package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"test/proto" //这里是proto生成的文件
)

const (
	address = "localhost:8089"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := proto.NewIrisServiceClient(conn)

	r, e := c.GetUserInfo(context.Background(), &proto.UserRequest{UserId: 1})

	if e != nil {
		log.Fatalf("could not greet: %v", e)
	}

	log.Println(r)
}

```
