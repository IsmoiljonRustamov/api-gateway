package service

import (
	"context"
	"database/sql"
	"fmt"

	"log"
	pc "projects/user-service/genproto/comment"
	pp "projects/user-service/genproto/post"
	pu "projects/user-service/genproto/user"
	l "projects/user-service/pkg/logger"
	grpcClient "projects/user-service/service/grpc_client"
	storage "projects/user-service/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	Client  grpcClient.Clients
}

func NewUserService(db *sql.DB, log l.Logger, client grpcClient.Clients) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		Client:  client,
	}
}

func (s *UserService) Create(ctx context.Context, req *pu.UserRequest) (*pu.UserResponse, error) {
	fmt.Println(req.UserName)
	user, err := s.storage.User().Create(&pu.UserRequest{
		Id:       req.Id,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		UserName: req.UserName,
	})

	if err != nil {
		s.logger.Error("error while creating")
		return &pu.UserResponse{}, err
	}
	return &pu.UserResponse{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: req.Password,
		UserName: req.UserName,
	}, nil
}

func (s *UserService) GetUserById(ctx context.Context, req *pu.IdRequest) (*pu.UserResponse, error) {
	user, err := s.storage.User().GetUserById(&pu.IdRequest{Id: req.Id})
	if err != nil {
		s.logger.Error("error while getting by id")
		return &pu.UserResponse{}, err
	}
	posts, err := s.Client.Post().GetAllPostsByUserId(ctx, &pp.IdRequest{Id: user.Id})
	if err != nil {
		log.Println("failed to getting user for getting user: ", err)
		return &pu.UserResponse{}, err
	}

	for _, post := range posts.Posts {
		comments, err := s.Client.Comment().GetCommentsForPost(ctx, &pc.GetAllCommentsRequest{PostId: post.Id})
		if err != nil {
			log.Println("failed to get comments for post from user-service: ", err)
			return &pu.UserResponse{}, err
		}

		for _, comment := range comments.Comments {
			comment.PostUserName = user.Name

			connUser, err := s.storage.User().GetUserById(&pu.IdRequest{Id: comment.UserId})
			if err != nil {
				log.Println("failed to get comment user: ", err)
				return &pu.UserResponse{}, err
			}
			comment.UserName = connUser.Name
		}
	}

	return user,nil
}

func (s *UserService) GetUserForClient(ctx context.Context, id *pu.IdRequest) (*pu.UserResponse, error) {
	user, err := s.storage.User().GetUserForClient(&pu.IdRequest{Id: id.Id})
	if err != nil {
		log.Println("failed to get user for client: ", err)
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUsers(ctx context.Context, req *pu.UserForGetUsers) (*pu.Users, error) {
	users, err := s.storage.User().GetUsers(req)
	if err != nil {
		log.Println("failed to get users indo: ", err)
		return &pu.Users{}, err
	}
	for _, user := range users.Users {

		posts, err := s.Client.Post().GetAllPostsByUserId(ctx, &pp.IdRequest{Id: user.Id})
		if err != nil {
			log.Println("failed to get post for gettig users: ", err)
			return &pu.Users{}, err
		}

		for _, post := range posts.Posts {
			var posT pu.Post
			posT.Title = post.Title
			posT.Description = post.Description

			comments, err := s.Client.Comment().GetCommentsForPost(ctx, &pc.GetAllCommentsRequest{PostId: post.Id})
			if err != nil {
				log.Println("failed to get comments for post from user-service: ", err)
				return &pu.Users{}, err
			}

			for _, comment := range comments.Comments {
				var commenT pu.Comments

				commenT.PostId = comment.PostId
				commenT.UserId = comment.UserId
				commenT.Text = comment.Text

				comment.PostUserName = user.Name
				comment.PostTitle = post.Title

				connUser, err := s.storage.User().GetUserById(&pu.IdRequest{Id: commenT.UserId})
				if err != nil {
					log.Println("failed to get comment user: ", err)
					return &pu.Users{}, err
				}
				comment.UserName = connUser.Name

				posT.Comments = append(posT.Comments, &commenT)

			}
			user.Posts = append(user.Posts, &posT)

		}
	}

	return users, nil

}

func (s *UserService) UpdateUser(ctx context.Context, user *pu.UserRequest) (*pu.UserForUpdate, error) {
	users, err := s.storage.User().UpdateUser(&pu.UserRequest{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		UserName: user.UserName,
	})
	if err != nil {
		log.Println("Failed to update user info: ", err)
	}

	return users,nil
}

func (s *UserService) DeleteUser(ctx context.Context, user *pu.IdRequest) (*pu.UserForUpdate, error) {
	_, err := s.storage.User().DeleteUser(user)
	if err != nil {
		log.Println("Failed to delete user info: ", err)
	}

	return &pu.UserForUpdate{}, nil
}

func (s *UserService) CheckField(ctx context.Context, req *pu.CheckFieldReq) (*pu.CheckFieldRes, error) {
	res, err := s.storage.User().CheckField(req)
	if err != nil {
		s.logger.Error("error check", l.Any("error check filed user", err))
		return &pu.CheckFieldRes{}, status.Error(codes.Internal, "something went wrong, please check user info")
	}

	return res, nil
}
