package config

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

func InitRpcServer(c *RpcConfig, router RpcRouter) error {
	d := NewRpcServer(router, AccLog)
	return d.Start(c)
}

func (c *RpcConfig) GetAddr() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}

func (s *RpcServer) Running() bool {
	select {
	case <-s.ctx.Done():
		return false
	default:
		return true
	}
}

func (s *RpcServer) Start(c *RpcConfig) error {
	listener, err := net.Listen("tcp", c.GetAddr())

	if err != nil {
		return err
	}

	s.Server = grpc.NewServer()
	s.router.RegRpcService(s)
	s.ctx, s.canceler = context.WithCancel(context.Background())

	go func() {
		err = s.Server.Serve(listener)

		if err != nil && s.Running() {
			s.Log.LogInfo("can't serve at <%s>", c.GetAddr())
		}
	}()

	return nil
}

func (s *RpcServer) Stop() {
	s.canceler()
	s.Server.Stop()
}

func GetRpc() *RpcConfig {
	return config.RpcConfig
}

func NewRpcServer(r RpcRouter, l *AccessLog) *RpcServer {
	return &RpcServer{router: r, Log: l, Server: grpc.NewServer()}
}
