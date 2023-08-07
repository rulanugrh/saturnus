package main

import (
	"fmt"
	"net"

	"github.com/rulanugrh/saturnus/configs"
	"github.com/rulanugrh/saturnus/helper"
	"github.com/rulanugrh/saturnus/pb"
	"github.com/rulanugrh/saturnus/services"
	"google.golang.org/grpc"
)

func main() {
	conf := configs.GetConfig()

	fmt.Println("Starting server on port :50051...")

	dsn := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	listener, err := net.Listen("tcp", dsn)
	if err != nil {
		helper.PrintError(err)
	}

	srv := &services.TodoNewService{}
	serv := grpc.NewServer()

	pb.RegisterTodoServiceServer(serv, srv)
	
	if errs := serv.Serve(listener); errs != nil {
		helper.PrintError(errs)
	}
}