package postgres

import (
	"database/sql"
	"fmt"
	pb "projects/user-service/genproto/user"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *pb.UserRequest) (*pb.UserResponse, error) {
	resp := pb.UserResponse{}
	err := r.db.QueryRow("INSERT INTO users(name,email,password,user_name) VALUES($1,$2,$3,$4) RETURNING id,name,email,password,user_name,created_at,updated_at", user.Name, user.Email,user.Password,user.UserName).Scan(
		&resp.Id, &resp.Name, &resp.Email,&resp.Password,&resp.UserName, &resp.CreatedAt, &resp.UpdatedAt,
	)
	if err != nil {
		return nil, err	
	}
	return &resp, nil
}

func (r *userRepo) GetUserById(user *pb.IdRequest) (*pb.UserResponse, error) {
	resp := pb.UserResponse{}
	err := r.db.QueryRow("SELECT id,name,email,password,user_name,created_at,updated_at FROM users WHERE id=$1", user.Id).Scan(
		&resp.Id, &resp.Name, &resp.Email, &resp.Password,&resp.UserName,&resp.CreatedAt, &resp.UpdatedAt,
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

func (r *userRepo) GetUsers(user *pb.UserForGetUsers) (*pb.Users,error) {
	var resp pb.Users
	offset := (user.Page - 1) * user.Limit

	rows,err := r.db.Query("SELECT id,name,email,created_at,updated_at FROM users LIMIT $1 OFFSET $2",user.Limit,offset)
	if err != nil {
		return nil,err
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
			fmt.Println("Failed to scan users info: ",err)
		}
		resp.Users = append(resp.Users, &resu)
	}

	return &resp,nil
}

func (r *userRepo) UpdateUser(user *pb.UserRequest) (*pb.UserForUpdate,error) {
	var res pb.UserForUpdate
	err := r.db.QueryRow("UPDATE users SET name=$1, email=$2,password=$3,user_name=$4 WHERE id=$5 RETURNING name,email,password,user_name,id",user.Name,user.Email,user.Password,user.UserName,user.Id).Scan(
		&res.Name,&res.Email,&res.Password,&res.UserName,&res.Id,
	)
	if err != nil {
		fmt.Println("Failed to update user info: ",err)
	}

	return &res,nil
}


func (r *userRepo) DeleteUser(user *pb.IdRequest) (*pb.UserForUpdate,error) {
	var resp pb.UserForUpdate
	err := r.db.QueryRow("DELETE FROM users WHERE id=$1 RETURNING id,name,email,password,user_name", user.Id).Scan(
		&resp.Id,&resp.Name,&resp.Email,&resp.Password,&resp.UserName,
	)
	if err != nil {
		fmt.Println("Failed to delete user info: ",err)
	}

	return &resp,nil
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
