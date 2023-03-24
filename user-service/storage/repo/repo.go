package repo

import (
	pb "projects/user-service/genproto/user"
)

type UserStorageI interface {
	Create(*pb.UserRequest) (*pb.UserResponse, error)
	GetUserById(*pb.IdRequest) (*pb.UserResponse, error)
	GetUserForClient(*pb.IdRequest) (*pb.UserResponse, error)
	GetUsers(*pb.UserForGetUsers) (*pb.Users,error)
	UpdateUser(*pb.UserRequest) (*pb.UserForUpdate,error)
	DeleteUser(*pb.IdRequest) (*pb.UserForUpdate,error) 
}
