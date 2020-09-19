package main

import (
	"github.com/atalia/grpc-tutorial/Server/route_guide"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main(){
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	// 注册服务
	route_guide.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Println(err)
	}
}
