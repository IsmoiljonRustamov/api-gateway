package grpcclient

import (
	"fmt"
	"projects/comment-service/config"
	p "projects/comment-service/genproto/post"
	u "projects/comment-service/genproto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)	

type Clients interface{
	User() u.UserServiceClient
	Post() p.PostServiceClient
}

type ServiceManagaer struct {
	Config config.Config
	userService u.UserServiceClient
	postService p.PostServiceClient
}

func New(cfg config.Config) (*ServiceManagaer,error) {
	connUser,err := grpc.Dial(
		fmt.Sprintf("%s:%s",cfg.UserServiceHost,cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("user service dial host:%s, port:%s",cfg.UserServiceHost,cfg.UserServicePort)
	}
	connPost,err := grpc.Dial(
		fmt.Sprintf("%s:%s",cfg.PostServiceHost,cfg.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil,fmt.Errorf("post service dial host:%s, port:%s",cfg.PostServiceHost,cfg.PostServicePort)
	}	

	return &ServiceManagaer{
		Config: cfg,
		userService: u.NewUserServiceClient(connUser),
		postService: p.NewPostServiceClient(connPost),
	},nil
}

func (s *ServiceManagaer) User() u.UserServiceClient {
	return s.userService
}

func (s *ServiceManagaer) Post() p.PostServiceClient {
	return s.postService
}