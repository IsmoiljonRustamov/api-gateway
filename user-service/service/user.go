package service

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"log"
	pu "projects/user-service/genproto/user"
	l "projects/user-service/pkg/logger"
	"projects/user-service/pkg/messagebroker"
	storage "projects/user-service/storage"
	"projects/user-service/storage/repo"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	producer map[string]messagebroker.Producer
}

func NewUserService(db *sql.DB, log l.Logger, producer map[string]messagebroker.Producer) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		producer: producer,
	}
}

func (s *UserService) produceMessage(raw *pu.PostForCreate) error {
	data,err := raw.Marshal()
	if err != nil {
		return err
	}
	logPost := raw.String()
	fmt.Println(logPost)
	err = s.producer["user"].Produce([]byte("user"), data, logPost)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Create(ctx context.Context, req *pu.UserRequest) (*pu.UserResponse, error) {
	user, err := s.storage.User().Create(repo.UserRequest{
		Id:           req.Id,
		Name:         req.Name,
		Email:        req.Email,
		UserType:     req.UserType,
		Password:     req.Password,
		UserName:     req.UserName,
		RefreshToken: req.RefreshToken,
	})

	if err != nil {
		s.logger.Error("error while creating")
		return &pu.UserResponse{}, err
	}
	post := req.Posts
	if post != nil {
		id := user.Id
		post.PosterId = id
		err =  s.produceMessage(post)
		if err != nil {
			fmt.Println(err)
		}
		return &pu.UserResponse{},nil
	} else {
		fmt.Println(user)
		return &pu.UserResponse{}, nil
	}

}

func (s *UserService) GetUserById(ctx context.Context, req *pu.IdRequest) (*pu.UserResponse, error) {
	user, err := s.storage.User().GetUserById(&pu.IdRequest{Id: req.Id})
	if err != nil {
		s.logger.Error("error while getting by id")
		return &pu.UserResponse{}, err
	}
	// posts, err := s.Client.Post().GetAllPostsByUserId(ctx, &pp.IdRequest{Id: user.Id})
	// if err != nil {
	// 	log.Println("failed to getting user for getting user: ", err)
	// 	return &pu.UserResponse{}, err
	// }

	// for _, post := range posts.Posts {
	// 	comments, err := s.Client.Comment().GetCommentsForPost(ctx, &pc.GetAllCommentsRequest{PostId: post.Id})
	// 	if err != nil {
	// 		log.Println("failed to get comments for post from user-service: ", err)
	// 		return &pu.UserResponse{}, err
	// 	}

	// 	for _, comment := range comments.Comments {
	// 		comment.PostUserName = user.Name

	// 		connUser, err := s.storage.User().GetUserById(&pu.IdRequest{Id: comment.UserId})
	// 		if err != nil {
	// 			log.Println("failed to get comment user: ", err)
	// 			return &pu.UserResponse{}, err
	// 		}
	// 		comment.UserName = connUser.Name
	// 	}
	// }

	return user, nil
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
	// for _, user := range users.Users {

	// 	posts, err := s.Client.Post().GetAllPostsByUserId(ctx, &pp.IdRequest{Id: user.Id})
	// 	if err != nil {
	// 		log.Println("failed to get post for gettig users: ", err)
	// 		return &pu.Users{}, err
	// 	}

	// 	for _, post := range posts.Posts {
	// 		var posT pu.Post
	// 		posT.Title = post.Title
	// 		posT.Description = post.Description

	// 		comments, err := s.Client.Comment().GetCommentsForPost(ctx, &pc.GetAllCommentsRequest{PostId: post.Id})
	// 		if err != nil {
	// 			log.Println("failed to get comments for post from user-service: ", err)
	// 			return &pu.Users{}, err
	// 		}

	// 		for _, comment := range comments.Comments {
	// 			var commenT pu.Comments

	// 			commenT.PostId = comment.PostId
	// 			commenT.UserId = comment.UserId
	// 			commenT.Text = comment.Text

	// 			comment.PostUserName = user.Name
	// 			comment.PostTitle = post.Title

	// 			connUser, err := s.storage.User().GetUserById(&pu.IdRequest{Id: commenT.UserId})
	// 			if err != nil {
	// 				log.Println("failed to get comment user: ", err)
	// 				return &pu.Users{}, err
	// 			}
	// 			comment.UserName = connUser.Name

	// 			posT.Comments = append(posT.Comments, &commenT)

	// 		}
	// 		user.Posts = append(user.Posts, &posT)

	// 	}
	// }

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

	return users, nil
}

func (s *UserService) DeleteUser(ctx context.Context, user *pu.IdRequest) (*pu.UserForUpdate, error) {
	_, err := s.storage.User().DeleteUser(user)
	if err != nil {
		log.Println("Failed to delete user info: ", err)
	}

	return &pu.UserForUpdate{}, nil
}

func (s *UserService) CheckField(ctx context.Context, req *pu.CheckFieldReq) (*pu.CheckFieldRes, error) {
	res, err := s.storage.User().CheckField(&pu.CheckFieldReq{
		Field: req.Field,
		Value: req.Value,
	})
	if err != nil {
		s.logger.Error("error check", l.Any("error check filed user", err))
		return &pu.CheckFieldRes{}, status.Error(codes.Internal, "something went wrong, please check user info")
	}

	return res, nil
}

func (s *UserService) Login(ctx context.Context, req *pu.LoginRequest) (*pu.LoginResponse, error) {
	req.Email = strings.ToLower(req.Email)
	req.Email = strings.TrimSpace(req.Email)

	user, err := s.storage.User().Login(req)
	if err != nil {
		s.logger.Error("error while getting user by email", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		s.logger.Error("error while comparing hashed password, Invalid credentials", l.Any("req", req))
		return nil, status.Error(codes.InvalidArgument, "Invalid credentials")
	}

	return user, nil
}

func (s *UserService) UpdateToken(ctx context.Context, req *pu.RequestForTokens) (*pu.LoginResponse, error) {
	res, err := s.storage.User().UpdateToken(req)
	if err != nil {
		s.logger.Error("error while updating tokens", l.Error(err), l.Any("req", req))
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return res, nil
}
