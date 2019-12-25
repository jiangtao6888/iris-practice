package routes

import (
	health "google.golang.org/grpc/health/grpc_health_v1"
	"iris/app/Rpc/Controllers"
	"iris/config"
	"iris/libraries/proto"
)

func (r *rpcRouter) RegRpcService(s *config.RpcServer) {
	proto.RegisterIrisServiceServer(s.Server, Controllers.User)
	health.RegisterHealthServer(s.Server, Controllers.Rpc)
}

var RpcRouter *rpcRouter

type rpcRouter struct{}
