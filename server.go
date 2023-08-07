package main

import (
	"fmt"
	"net"

	"github.com/rulanugrh/saturnus/configs"
	"github.com/rulanugrh/saturnus/helper"
	"github.com/rulanugrh/saturnus/pb"
	"github.com/rulanugrh/saturnus/repository"
	"github.com/rulanugrh/saturnus/services"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func main() {
	conf := configs.GetConfig()

	fmt.Printf("Starting server on port %s:%s", conf.Host, conf.Port)

	dsn := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	listener, err := net.Listen("tcp", dsn)
	if err != nil {
		helper.PrintError(err)
	}

	var todoColl *mongo.Collection = configs.MongoColl("todo", configs.DB)
	todoRepository := repository.TodoRepository(todoColl)
	srv := services.NewGrpcPost(todoRepository, todoColl)
	serv := grpc.NewServer()

	pb.RegisterTodoServiceServer(serv, srv)
	
	if errs := serv.Serve(listener); errs != nil {
		helper.PrintError(errs)
	}
}