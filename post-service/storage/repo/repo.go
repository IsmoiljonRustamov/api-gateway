package repo

import (
	p "projects/post-service/genproto/post"
)

type PostStorageI interface {
	Create(*p.PostRequest) (*p.PostResponse, error)
	GetPostById(*p.IdRequest) (*p.PostResponse,error)
	GetAllPostsByUserId(*p.IdRequest) (*p.Posts,error)
	GetPostForComment(*p.IdRequest) (*p.PostResponse, error)
	GetPostForUser(*p.IdRequest) (*p.Posts, error)
	GetPosts(*p.GetForPosts) (*p.Posts,error)
	UpdatePost(*p.RequestForUpdate) (*p.PostResponse,error)
	DeletePost(*p.IdRequest) (*p.PostResponse,error)
}

