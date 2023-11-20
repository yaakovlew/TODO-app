package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"task/pkg/grpc_handler/api"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	task := api.TaskInput{
		IsDone:      false,
		Header:      "Check",
		Description: "Only for check",
		Date:        "2023-11-16T10:11:18Z",
	}

	handler := api.NewTaskClient(conn)
	if _z, err := handler.CreateTask(context.Background(), &task); err != nil {
		log.Fatal(err)
	}
}
