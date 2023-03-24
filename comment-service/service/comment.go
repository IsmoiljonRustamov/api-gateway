package service

import (
	"context"
	"database/sql"
	"log"
	c "projects/comment-service/genproto/comment"
	p "projects/comment-service/genproto/post"
	u "projects/comment-service/genproto/user"
	"projects/comment-service/pkg/logger"
	grpc_client "projects/comment-service/service/grpc_client"

	"projects/comment-service/storage"
)

type CommentService struct {
	storage storage.IStorage
	Logger  logger.Logger
	Client  grpc_client.Clients
}

func NewCommentService(db *sql.DB, log logger.Logger, client grpc_client.Clients) *CommentService {
	return &CommentService{
		storage: storage.NewStoragePG(db),
		Logger:  log,
		Client:  client,
	}
}

func (s *CommentService) CreateComment(ctx context.Context, req *c.CommentRequest) (*c.CommentResponse, error) {
	res, err := s.storage.Comment().CreateComment(req)
	if err != nil {
		log.Println("failed to create comment: ", err)
		return &c.CommentResponse{}, err
	}

	post, err := s.Client.Post().GetPostForComment(ctx, &p.IdRequest{Id: req.PostId})
	if err != nil {
		log.Println("failed to getting post for create comment: ", err)
		return &c.CommentResponse{}, err
	}
	res.PostTitle = post.Title

	user, err := s.Client.User().GetUserForClient(ctx, &u.IdRequest{Id: req.UserId})
	if err != nil {
		log.Println("failed to getting user for create comment: ", err)
		return &c.CommentResponse{}, err
	}
	res.UserName = user.Name

	postUser, err := s.Client.User().GetUserForClient(ctx, &u.IdRequest{Id: post.UserId})
	if err != nil {
		log.Println("failed to getting post's user for create comment: ", err)
		return &c.CommentResponse{}, err
	}
	res.PostUserName = postUser.Name

	return res, nil

}

func (s *CommentService) GetComments(ctx context.Context, req *c.ForGetComments) (*c.Comments, error) {
	comments, err := s.storage.Comment().GetComments(req)
	if err != nil {
		log.Println("Failed to get comments: ", err)
		return &c.Comments{}, err
	}
	for _, comment := range comments.Comments {
		user, err := s.Client.User().GetUserForClient(ctx, &u.IdRequest{Id: comment.UserId})
		if err != nil {
			return nil, err
		}
		comment.UserName = user.Name
		comment.UserId = user.Id

		posts, err := s.Client.Post().GetPostForComment(ctx, &p.IdRequest{Id: comment.PostId})
		if err != nil {
			return nil, err
		}

		comment.PostId = posts.Id
		comment.PostTitle = posts.Title

		connUser, err := s.Client.User().GetUserForClient(ctx, &u.IdRequest{Id: posts.UserId})
		if err != nil {
			return nil, err
		}
		comment.PostUserName = connUser.Name
	}

	return comments, nil
}

func (s *CommentService) GetCommentsForPost(ctx context.Context, req *c.GetAllCommentsRequest) (*c.Comments, error) {
	res, err := s.storage.Comment().GetComment(req)
	if err != nil {
		log.Println("Failed to get comments for client: ", err)
		return &c.Comments{}, err
	}

	return res, nil
}

func (s *CommentService) DeleteComment(ctx context.Context, req *c.IdRequest) (*c.CommentResponse, error) {
	res, err := s.storage.Comment().DeleteComment(req)
	if err != nil {
		log.Println("Failed to delete comment: ", err)
		return &c.CommentResponse{}, err
	}

	post, err := s.Client.Post().GetPostForComment(ctx, &p.IdRequest{Id: res.PostId})
	if err != nil {
		log.Println("Failed to get post for delete comment: ", err)
		return &c.CommentResponse{}, err
	}
	res.PostTitle = post.Title

	user, err := s.Client.User().GetUserForClient(ctx, &u.IdRequest{Id: res.UserId})
	if err != nil {
		log.Println("Failed to get user for delete comment: ", err)
		return &c.CommentResponse{}, err
	}
	res.UserName = user.Name

	postUser, err := s.Client.User().GetUserForClient(ctx, &u.IdRequest{Id: post.UserId})
	if err != nil {
		log.Println("Failed to get post's user for delete comment: ", err)
		return &c.CommentResponse{}, err
	}
	res.PostUserName = postUser.Name

	return res, nil

}

func (s *CommentService) UpdateComment(ctx context.Context, req *c.ForUpdate) (*c.CommentResponse, error) {
	res, err := s.storage.Comment().UpdateComment(req)
	if err != nil {
		log.Println("Failed to update comment from comment service: ", err)
		return &c.CommentResponse{}, err
	}

	post, err := s.Client.Post().GetPostForComment(ctx, &p.IdRequest{Id: res.PostId})
	if err != nil {
		log.Println("Failed to get post for update comment: ", err)
		return &c.CommentResponse{}, err
	}
	res.PostId = post.Id
	res.PostTitle = post.Title

	user, err := s.Client.User().GetUserForClient(ctx, &u.IdRequest{Id: res.UserId})
	if err != nil {
		log.Println("Failed to get user for update comment: ", err)
		return &c.CommentResponse{}, err
	}
	res.UserId = user.Id
	res.UserName = user.Name

	postUser, err := s.Client.User().GetUserForClient(ctx, &u.IdRequest{Id: post.UserId})
	if err != nil {
		log.Println("Failed to get post's user for update comment: ", err)
		return &c.CommentResponse{}, err
	}
	res.PostUserName = postUser.Name

	return res, nil
}

func (s *CommentService) GetComment(ctx context.Context, req *c.GetAllCommentsRequest) (*c.Comments, error) {
	comment, err := s.storage.Comment().GetComment(&c.GetAllCommentsRequest{PostId: req.PostId})
	if err != nil {
		return nil, err
	}

	return comment, nil
}
