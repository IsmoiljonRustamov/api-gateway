package grpcclient

import (
	"fmt"
	"projects/post-service/config"
	pb "projects/post-service/genproto/user"
	pu "projects/post-service/genproto/comment"


	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients interface{
	User() pb.UserServiceClient
	Comment() pu.CommentServiceClient
}

type ServiceManager struct {
	Config config.Config
	userService pb.UserServiceClient
	commentService pu.CommentServiceClient
}

func New(cfg config.Config) (*ServiceManager,error){
	connUser,err := grpc.Dial(
		fmt.Sprintf("%s:%s",cfg.UserServiceHost,cfg.UserServicePort ),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil,fmt.Errorf("user service dial host:%s port:%s", cfg.UserServiceHost,cfg.UserServicePort)
	}

	connComment,err := grpc.Dial(
		fmt.Sprintf("%s:%s",cfg.CommentServiceHost,cfg.CommentServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil,fmt.Errorf("comment service dial host:%s port=%s",cfg.CommentServiceHost,cfg.CommentServicePort)
	}	

	return &ServiceManager{
		Config: cfg,
		userService: pb.NewUserServiceClient(connUser),
		commentService: pu.NewCommentServiceClient(connComment),
	},nil

}

func (s *ServiceManager) User() pb.UserServiceClient {
	return s.userService
}
func (s *ServiceManager) Comment() pu.CommentServiceClient {
	return s.commentService
}