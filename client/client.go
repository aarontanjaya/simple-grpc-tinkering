package client

import (
	"grpc-crash-course/config"
	"log"

	pb "grpc-crash-course/todoProto"

	"google.golang.org/grpc"
)

func ClientTodo() pb.TodoClient {
	conn, err := grpc.Dial(config.SERVICE_PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatal("error connecting to service on port", config.SERVICE_PORT)
	}

	return pb.NewTodoClient(conn)
}
