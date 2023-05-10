package main

import (
	"net"
	"projects/user-service/config"
	pb "projects/user-service/genproto/user"
	"projects/user-service/kafka"
	"projects/user-service/pkg/db"
	"projects/user-service/pkg/logger"
	"projects/user-service/pkg/messagebroker"
	"projects/user-service/service"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "user-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlConfig", logger.String("host", cfg.PostgresHost), logger.String("port", cfg.PostgresPort), logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatal("connection to postgres error", logger.Error(err))
	}
	produceMap := make(map[string]messagebroker.Producer)
	topic := "topic1"
	userTopicProduce := kafka.NewKafkaProducer(cfg, log, topic)
	defer func() {
		err := userTopicProduce.Stop()
		if err != nil {
			log.Fatal("Failed to stopping Kafka", logger.Error(err))
		}
	}()
	produceMap["user"] = userTopicProduce

	userService := service.NewUserService(connDB, log, produceMap)

	lis, err := net.Listen("tcp", cfg.UserServicePort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterUserServiceServer(s, userService)
	log.Info("main: server running",
		logger.String("port", cfg.UserServicePort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

}
