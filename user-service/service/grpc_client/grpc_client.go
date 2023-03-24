package grpcclient

import (
	"fmt"
	"projects/user-service/config"
	pb "projects/user-service/genproto/post"
	pu "projects/user-service/genproto/comment"


	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients interface {
	Post() pb.PostServiceClient
	Comment() pu.CommentServiceClient

}

type ServiceManager struct {
	Config      config.Config
	postService pb.PostServiceClient
	commentService pu.CommentServiceClient
}

func New(cfg config.Config) (*ServiceManager, error) {
	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("post service dial host:%s port:%s", cfg.PostServiceHost, cfg.PostServicePort)
	}

	connComment,err := grpc.Dial(
		fmt.Sprintf("%s:%s",cfg.CommentServiceHost,cfg.CommentServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil,fmt.Errorf("comment service dial host:%s port:%s ",cfg.CommentServiceHost,cfg.CommentServicePort)
	}	

	return &ServiceManager{
		Config:      cfg,
		postService: pb.NewPostServiceClient(connPost),
		commentService: pu.NewCommentServiceClient(connComment),
	}, nil
}

func (s *ServiceManager) Post() pb.PostServiceClient {
	return s.postService
}

func (s *ServiceManager) Comment() pu.CommentServiceClient {
	return s.commentService
}
