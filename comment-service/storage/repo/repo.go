package repo

import (
	c "projects/comment-service/genproto/comment"

)

type CommentStorageI interface {
	CreateComment(*c.CommentRequest) (*c.CommentResponse,error) 
	GetComments(*c.ForGetComments) (*c.Comments,error)
	GetCommentsForPost(*c.GetAllCommentsRequest) (*c.Comments,error)
	DeleteComment(*c.IdRequest) (*c.CommentResponse,error)
	UpdateComment(*c.ForUpdate) (*c.CommentResponse,error)
	GetComment(*c.GetAllCommentsRequest) (*c.Comments,error)
}

