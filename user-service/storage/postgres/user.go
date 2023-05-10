package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	pb "projects/user-service/genproto/user"
	"projects/user-service/storage/repo"
)

type userRepo struct {
	db *sql.DB
}

var (
	invalidArgumentError = errors.New("invalid argment")
)

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user repo.UserRequest) (repo.UserResponse, error) {
	resp := pb.UserResponse{}
	err := r.db.QueryRow("INSERT INTO users(id,name,email,user_type,password,user_name,refresh_token) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id,name,email,password,user_name,user_name,created_at,updated_at,refresh_token", user.Id, user.Name, user.Email, user.UserType, user.Password, user.UserName, user.RefreshToken).Scan(
		&resp.Id, &resp.Name, &resp.Email, &resp.Password, &resp.UserType, &resp.UserName, &resp.CreatedAt, &resp.UpdatedAt, &resp.RefreshToken,
	)
	if err != nil {	
		return repo.UserResponse{}, err
	}
	return repo.UserResponse{}, nil
}

func (r *userRepo) GetUserById(user *pb.IdRequest) (*pb.UserResponse, error) {
	resp := pb.UserResponse{}
	err := r.db.QueryRow("SELECT id,name,email,password,user_name,created_at,updated_at FROM users WHERE id=$1", user.Id).Scan(
		&resp.Id, &resp.Name, &resp.Email, &resp.Password, &resp.UserName, &resp.CreatedAt, &resp.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (r *userRepo) GetUserForClient(id *pb.IdRequest) (*pb.UserResponse, error) {
	resp := pb.UserResponse{}
	err := r.db.QueryRow("SELECT id,name,email,created_at,updated_at FROM users WHERE id=$1", id.Id).Scan(
		&resp.Id, &resp.Name, &resp.Email, &resp.CreatedAt, &resp.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil

}

func (r *userRepo) GetUsers(user *pb.UserForGetUsers) (*pb.Users, error) {
	var resp pb.Users
	offset := (user.Page - 1) * user.Limit

	rows, err := r.db.Query("SELECT id,name,email,created_at,updated_at FROM users LIMIT $1 OFFSET $2", user.Limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var resu pb.UserResponse
		err := rows.Scan(
			&resu.Id,
			&resu.Name,
			&resu.Email,
			&resu.CreatedAt,
			&resu.UpdatedAt,
		)
		if err != nil {
			fmt.Println("Failed to scan users info: ", err)
		}
		resp.Users = append(resp.Users, &resu)
	}

	return &resp, nil
}

func (r *userRepo) UpdateUser(user *pb.UserRequest) (*pb.UserForUpdate, error) {
	var res pb.UserForUpdate
	err := r.db.QueryRow("UPDATE users SET name=$1, email=$2,password=$3,user_name=$4 WHERE id=$5 RETURNING name,email,password,user_name,id", user.Name, user.Email, user.Password, user.UserName, user.Id).Scan(
		&res.Name, &res.Email, &res.Password, &res.UserName, &res.Id,
	)
	if err != nil {
		fmt.Println("Failed to update user info: ", err)
	}

	return &res, nil
}

func (r *userRepo) DeleteUser(user *pb.IdRequest) (*pb.UserForUpdate, error) {
	var resp pb.UserForUpdate
	err := r.db.QueryRow("DELETE FROM users WHERE id=$1 RETURNING id,name,email,password,user_name", user.Id).Scan(
		&resp.Id, &resp.Name, &resp.Email, &resp.Password, &resp.UserName,
	)
	if err != nil {
		fmt.Println("Failed to delete user info: ", err)
	}

	return &resp, nil
}

func (r *userRepo) CheckField(req *pb.CheckFieldReq) (*pb.CheckFieldRes, error) {
	query := fmt.Sprintf("SELECT 1 FROM users WHERE %s=$1", req.Field)
	var temp int
	err := r.db.QueryRow(query, req.Value).Scan(&temp)
	if err == sql.ErrNoRows {
		return &pb.CheckFieldRes{Exists: false}, nil
	}

	if err != nil {
		return nil, err
	}

	if temp == 0 {
		return &pb.CheckFieldRes{Exists: true}, nil
	}

	return &pb.CheckFieldRes{Exists: false}, nil
}

func (r *userRepo) Login(req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var res pb.LoginResponse
	err := r.db.QueryRow("SELECT id,name,email,password,user_name,refresh_token, user_type FROM users WHERE email=$1", req.Email).Scan(
		&res.Id, &res.Name, &res.Email, &res.Password, &res.UserName,&res.RefreshToken,&res.UserType,
	)
	if err != nil {
		log.Println("Failed to select user info from login func in psql: ", err)
	}

	return &res, nil
}

func (r *userRepo) UpdateToken(req *pb.RequestForTokens) (*pb.LoginResponse, error) {
	var resp pb.LoginResponse
	err := r.db.QueryRow("UPDATE users SET refresh_token=$1 WHERE id=$2 RETURNING id,name,email,password,user_name,refresh_token",req.RefreshToken, req.Id).Scan(
		&resp.Id, &resp.Name, &resp.Email, &resp.Password, &resp.UserName, &resp.RefreshToken,
	)
	if err != nil {
		log.Println("Failed to update tokens: ", err)
	}

	return &resp, nil
}
