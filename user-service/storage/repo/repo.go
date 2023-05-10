package repo

import (
	pb "projects/user-service/genproto/user"
)

type UserStorageI interface {
	Create(UserRequest) (UserResponse, error)
	GetUserById(*pb.IdRequest) (*pb.UserResponse, error)
	GetUserForClient(*pb.IdRequest) (*pb.UserResponse, error)
	GetUsers(*pb.UserForGetUsers) (*pb.Users,error)
	UpdateUser(*pb.UserRequest) (*pb.UserForUpdate,error)
	DeleteUser(*pb.IdRequest) (*pb.UserForUpdate,error) 
	CheckField(*pb.CheckFieldReq) (*pb.CheckFieldRes, error)
	Login(*pb.LoginRequest) (*pb.LoginResponse,error)
	UpdateToken(*pb.RequestForTokens) (*pb.LoginResponse,error)
}

type UserRequest struct {
	Id string
	Name string
	Email string
	UserType string 
	Password string
	UserName string
	RefreshToken string
}

type UserResponse struct {
	Id string
	Name string
	Email string
	Password string
	UserName string
	RefreshToken string
	AccesToken string
	CreatedAt string
	UpdatedAt string
}