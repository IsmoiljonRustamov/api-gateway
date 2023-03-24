package main

import (
	"fmt"
	"net"
	"projects/comment-service/config"
	c "projects/comment-service/genproto/comment"
	"projects/comment-service/pkg/db"
	"projects/comment-service/pkg/logger"
	"projects/comment-service/service"
	grpcClient "projects/comment-service/service/grpc_client"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "comment-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlConfig", logger.String("host", cfg.PostgresHost), logger.String("port", cfg.PostgresPort), logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatal("connection to postgres error", logger.Error(err))
	}

	grpc_client, err := grpcClient.New(cfg)
	if err != nil {
		fmt.Println("error while grpc client from comment-service", err.Error())
	}

	commentService := service.NewCommentService(connDB, log, grpc_client)

	lis, err := net.Listen("tcp", cfg.CommentServicePort)
	if err != nil {
		log.Fatal("error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	reflection.Register(s)
	c.RegisterCommentServiceServer(s, commentService)

	log.Info("main: server running",
		logger.String("port", cfg.CommentServicePort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening:%v", logger.Error(err))
	}

}
