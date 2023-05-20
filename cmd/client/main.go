package main

import (
	"context"
	"grpc-crash-course/client"
	pb "grpc-crash-course/todoProto"
	"io"
	"log"
)

func main() {
	client := client.ClientTodo()

	todos := []string{
		"Makan",
		"Minum",
		"Tidur",
		"Baca",
	}

	for i, act := range todos {
		item, err := client.CreateTodo(context.Background(), &pb.TodoItem{
			Text: act,
			Id:   int32(i),
		})
		if err != nil {
			log.Fatal("error: ", err)
		}

		log.Printf("item created: %v \n", item)
	}

	list, err := client.ReadTodos(context.Background(), &pb.Void{})
	if err != nil {
		log.Fatal("error: ", err)
	}

	log.Printf("list of items: %v", list)

	stream, err := client.ReadTodosStream(context.Background(), &pb.Void{})
	if err != nil {
		log.Fatal("error: ", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Fatal("failed streaming")
			}
			log.Printf("Stream result : %v", resp)
		}
	}()

	<-done
	log.Println("Done")
}
