package storage

import (
	"database/sql"
	"projects/user-service/storage/repo"

	"projects/user-service/storage/postgres"
)

type IStorage interface {
	User() repo.UserStorageI
}

type storagePG struct {
	db       *sql.DB
	userRepo repo.UserStorageI
}

func NewStoragePg(db *sql.DB) *storagePG {
	return &storagePG{
		db:       db,
		userRepo: postgres.NewUserRepo(db),
	}
}

func (s storagePG) User() repo.UserStorageI {
	return s.userRepo
}
