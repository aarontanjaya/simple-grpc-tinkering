package main

import (
	"context"
	"grpc-crash-course/client"
	pb "grpc-crash-course/todoProto"
	"log"
)

func main() {
	client := client.ClientTodo()
	item, err := client.CreateTodo(context.Background(), &pb.TodoItem{
		Text: "Makan",
		Id:   1,
	})
	if err != nil {
		log.Fatal("error: ", err)
	}

	log.Printf("item created: %v \n", item)

	list, err := client.ReadTodos(context.Background(), &pb.Void{})
	if err != nil {
		log.Fatal("error: ", err)
	}

	log.Printf("list of items: %v", list)
}
