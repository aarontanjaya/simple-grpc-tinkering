package server

import (
	"context"
	pb "grpc-crash-course/todoProto"
	"log"
)

type UserServer struct{}

var (
	userServer   UserServer
	localStorage *pb.TodoItems
)

func init() {
	localStorage = new(pb.TodoItems)
	localStorage.Items = make([]*pb.TodoItem, 0)
}

func (UserServer) CreateTodo(ctx context.Context, item *pb.TodoItem) (*pb.TodoItem, error) {
	localStorage.Items = append(localStorage.Items, item)
	log.Printf("todo item created %v \n", item)
	return item, nil
}

func (UserServer) ReadTodos(ctx context.Context, void *pb.Void) (*pb.TodoItems, error) {
	log.Println("read todos called")
	return localStorage, nil
}
