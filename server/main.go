package main

import (
	"API-Gateway/api"
	"API-Gateway/api/handler"
	"API-Gateway/genproto/content"
	"API-Gateway/genproto/users"
	"API-Gateway/logger"
	"log"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var loge *zap.Logger

func initLog() {
	log, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}
	loge = log
}

func main() {
	initLog()
	hand := NewHandler()
	router := api.Router(hand)
	log.Printf("server is running...")
	log.Fatal(router.Run(":8080"))
}

func NewHandler() *handler.Handler {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	con, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	contenservice := content.NewContentClient(conn)
	userservice := users.NewUserServiceClient(con)
	return &handler.Handler{ContentService: contenservice, Log: loge, UserService: userservice}
}
