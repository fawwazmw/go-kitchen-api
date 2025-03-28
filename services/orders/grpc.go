package main

import (
	handler "github.com/fawwazmw/go-kitchen-api/services/orders/handler/orders"
	"github.com/fawwazmw/go-kitchen-api/services/orders/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type gRPCServer struct {
	addr string
}

func newGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	grcpServer := grpc.NewServer()

	// register our grpc services
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grcpServer, orderService)

	log.Println("Starting gRPC server on", s.addr)

	return grcpServer.Serve(lis)
}
