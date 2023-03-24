package main

import (
	"fmt"
	"net"
	"projects/post-service/config"
	pb "projects/post-service/genproto/post"
	"projects/post-service/pkg/db"
	"projects/post-service/pkg/logger"
	"projects/post-service/service"
	grpc_client "projects/post-service/service/grpc_client"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "post-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlConfig", logger.String("host: ", cfg.PostServiceHost), logger.String("port: ", cfg.PostServicePort), logger.String("database: ", cfg.PostgresDatabase))

	connDB, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatal("connection to postgres error", logger.Error(err))
	}

	grpcclient,err := grpc_client.New(cfg)
	if err != nil {
		fmt.Println("error while grpc client from post-service",err.Error())
	}
	postService := service.NewPostService(connDB, log,grpcclient)

	lis, err := net.Listen("tcp", cfg.PostServicePort)
	if err != nil {
		log.Fatal("error while listening: %v", logger.Error(err))
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterPostServiceServer(s, postService)
	log.Info("main: server running",
		logger.String("port", cfg.PostServicePort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

}
