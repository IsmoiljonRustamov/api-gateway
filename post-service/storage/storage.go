package storage

import (
	"database/sql"

	"projects/post-service/storage/postgres"
	"projects/post-service/storage/repo"
)



type IStorage interface {
	Post() repo.PostStorageI
}


type storagePG struct {
	db *sql.DB
	postRepo repo.PostStorageI
}

func NewStoragePg(db *sql.DB) *storagePG {
	return &storagePG{
		db: db,
		postRepo: postgres.NewPostRepo(db),
	}
}

func (s storagePG) Post() repo.PostStorageI {
	return s.postRepo
}