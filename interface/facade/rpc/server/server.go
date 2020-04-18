package server

import (
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"ddd/interface/dto/rpc"
	"ddd/interface/facade/rpc/handler"
)

const (
	address = "localhost:9090"
)

// Run 启动grpc server
func Run() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	s := grpc.NewServer()
	rpc.RegisterLeaverServer(s, &handler.CreateLeave{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("%+v", err)
	}
}
