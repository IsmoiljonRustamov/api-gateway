package storage

import (
	"database/sql"
	"projects/comment-service/storage/postgres"
	"projects/comment-service/storage/repo"
)

type IStorage interface {
	Comment() repo.CommentStorageI
}

type StoragePG struct {
	db          *sql.DB
	commentRepo repo.CommentStorageI
}

func NewStoragePG(db *sql.DB) *StoragePG {
	return &StoragePG{
		db:          db,
		commentRepo: postgres.NewCommentRepo(db),
	}
}

func (s StoragePG) Comment() repo.CommentStorageI {
	return s.commentRepo
}
