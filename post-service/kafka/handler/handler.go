package kafka

import (
	"fmt"
	"projects/post-service/config"
	"projects/post-service/pkg/logger"
	"projects/post-service/storage"

	pb "projects/post-service/genproto/post"
)


type KafkaHandler struct {
	config config.Config
	storage storage.IStorage
	log logger.Logger
}

func NewKafkaHandlerFunc(config config.Config, storage storage.IStorage, log logger.Logger) *KafkaHandler {
	return &KafkaHandler{
		config: config,
		storage: storage,
		log: log,
	}
}

func (h *KafkaHandler) Handle(value []byte) error {
	post := pb.PostRequest{}
	err := post.Unmarshal(value)
	if err != nil {
		return  err
	}
	fmt.Println("aaaaaaaaa")
	fmt.Println(post)
	_, err = h.storage.Post().Create(&pb.PostRequest{
		UserId: post.UserId,
		Description: post.Description,
	})
	if err != nil {
		return err
	}
	return nil
}