package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rulanugrh/saturnus/configs"
	"github.com/rulanugrh/saturnus/helper"
	"github.com/rulanugrh/saturnus/pb"
	"github.com/rulanugrh/saturnus/repository"
	"github.com/rulanugrh/saturnus/services"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conf := configs.GetConfig()

	dsnGRPC := fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.GRPCPort)
	dsnHTTP := fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.HTTPort)
	listener, err := net.Listen("tcp", dsnGRPC)
	if err != nil {
		helper.PrintError(err)
	}

	// Create COLLECTION and repository for services
	var todoColl *mongo.Collection = configs.MongoColl("todo", configs.DB)
	todoRepository := repository.TodoRepository(todoColl)
	srv := services.NewGrpcPost(todoRepository, todoColl)

	// Starting grpc serevr
	serv := grpc.NewServer()
	pb.RegisterTodoServiceServer(serv, srv)

	go func() {
		log.Fatalln(serv.Serve(listener))
	}()

	log.Printf("GRPC Server running on %s:%s", conf.Server.Host, conf.Server.GRPCPort)
	// create grpc connection for http
	maxMsgSize := 1024 * 1024 * 20
	connGrpc, errCon := grpc.DialContext(
		context.Background(),
		dsnGRPC,
		grpc.WithBlock(), 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize), grpc.MaxCallSendMsgSize(maxMsgSize)),
	)
	if errCon != nil {
		helper.PrintError(errCon)
	}
	
	// server for running http server
	srvHttp := runtime.NewServeMux()
	errHttp := pb.RegisterTodoServiceHandler(context.Background(), srvHttp, connGrpc)
	if errHttp != nil {
		helper.PrintError(errHttp)
	}


	srvHttpServ := &http.Server {
		Addr: dsnHTTP,
		Handler: srvHttp,
	}

	
	log.Printf("HTTP Server running on %s:%s", conf.Server.Host, conf.Server.HTTPort)
	if errServHttp := srvHttpServ.ListenAndServe(); errServHttp != nil {
		helper.PrintError(errServHttp)
	}
	
}