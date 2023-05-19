package server

import (
	"grpc-crash-course/config"
	pb "grpc-crash-course/todoProto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	server *grpc.Server
)

func InitServer() {
	server = grpc.NewServer()
	pb.RegisterTodoServer(server, userServer)

	log.Println("starting rpc server")
	lis, err := net.Listen("tcp", config.SERVICE_PORT)
	if err != nil {
		log.Println("err:", err)
	}

	log.Fatal(server.Serve(lis))

}
