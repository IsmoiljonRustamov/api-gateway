package service

import (
	"context"
	"database/sql"
	"log"
	pc "projects/post-service/genproto/comment"
	pb "projects/post-service/genproto/post"
	pu "projects/post-service/genproto/user"
	l "projects/post-service/pkg/logger"
	grpc_client "projects/post-service/service/grpc_client"
	"projects/post-service/storage"
)

type PostService struct {
	storage storage.IStorage
	logger  l.Logger
	Client  grpc_client.Clients
}

func NewPostService(db *sql.DB, log l.Logger, client grpc_client.Clients) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		Client:  client,
	}
}

func (s *PostService) Create(ctx context.Context, req *pb.PostRequest) (*pb.PostResponse, error) {
	post, err := s.storage.Post().Create(req)
	if err != nil {
		s.logger.Error("error while created from post-service")
		return nil, err
	}
	user, err := s.Client.User().GetUserForClient(ctx, &pu.IdRequest{Id: req.UserId})
	if err != nil {
		log.Println("failed to getting user for create post: ", err)
		return &pb.PostResponse{}, err
	}

	post.UserName = user.Name
	post.UserEmail = user.Email

	return post, nil
}

func (s *PostService) GetPostById(ctx context.Context, req *pb.IdRequest) (*pb.PostResponse, error) {
	post, err := s.storage.Post().GetPostById(req)
	if err != nil {
		s.logger.Error("error while created from post-service")
		return nil, err
	}

	user, err := s.Client.User().GetUserForClient(ctx, &pu.IdRequest{Id: post.UserId})
	if err != nil {
		log.Println("failed to getting user for getting post: ", err)
		return &pb.PostResponse{}, err
	}
	post.UserName = user.Name
	post.UserEmail = user.Email

	comments, err := s.Client.Comment().GetCommentsForPost(ctx, &pc.GetAllCommentsRequest{PostId: post.Id})
	if err != nil {
		log.Println("failed to get comment for get post")
		return &pb.PostResponse{}, err
	}

	for _, comment := range comments.Comments {
		comUser, err := s.Client.User().GetUserForClient(ctx, &pu.IdRequest{Id: comment.UserId})
		if err != nil {
			log.Println("failed to get user for comment", err)
			return nil, err
		}
		comment.PostTitle = post.Title
		comment.UserId = comUser.Id
		comment.UserName = comUser.Name
		comment.PostUserName = user.Name
	}

	return post, nil
}

func (s *PostService) GetAllPostsByUserId(ctx context.Context, req *pb.IdRequest) (*pb.Posts, error) {
	posts, err := s.storage.Post().GetAllPostsByUserId(&pb.IdRequest{
		Id: req.Id,
	})
	if err != nil {
		s.logger.Error("error while getting all for getting post:")
		return nil, err
	}

	user, err := s.Client.User().GetUserForClient(ctx, &pu.IdRequest{Id: req.Id})
	if err != nil {
		log.Println("failed to get user for post: ", err)
		return nil, err
	}

	for _, post := range posts.Posts {
		post.UserName = user.Name
		post.UserEmail = user.Email

		comments, err := s.Client.Comment().GetCommentsForPost(ctx, &pc.GetAllCommentsRequest{PostId: post.Id})
		if err != nil {
			log.Println("failed to get comment for post", err)
			return nil, err
		}

		for _, comment := range comments.Comments {
			comUser, err := s.Client.User().GetUserForClient(ctx, &pu.IdRequest{Id: comment.UserId})
			if err != nil {
				log.Println("failed to get user for get comment: ", err)
				return nil, err
			}
			comment.PostTitle = post.Title
			comment.UserName = comUser.Name
			comment.PostUserName = user.Name
		}
	}

	return posts, nil
}

func (s *PostService) GetPostForComment(ctx context.Context, id *pb.IdRequest) (*pb.PostResponse, error) {
	post, err := s.storage.Post().GetPostForComment(id)
	if err != nil {
		log.Println("failed to get post: ", err)
		return nil, err
	}

	return post, nil
}

func (s *PostService) GetPostForUser(ctx context.Context, id *pb.IdRequest) (*pb.Posts, error) {
	posts, err := s.storage.Post().GetPostForUser(id)
	if err != nil {
		log.Println("failed to get posts for user: ", err)
		return nil, err
	}

	return posts, nil
}

func (s *PostService) GetPosts(ctx context.Context, post *pb.GetForPosts) (*pb.Posts, error) {
	posts, err := s.storage.Post().GetPosts(post)
	if err != nil {
		log.Println("Failed to get posts with limit and page: ", err)
	}

	for _, post := range posts.Posts {
		user, err := s.Client.User().GetUserForClient(ctx, &pu.IdRequest{Id: post.UserId})
		if err != nil {
			return nil, err
		}
		post.UserName = user.Name
		post.UserId = user.Id
		post.UserEmail = user.Email
		comments, err := s.Client.Comment().GetCommentsForPost(ctx, &pc.GetAllCommentsRequest{PostId: post.Id})
		if err != nil {
			return nil, err
		}
		for _, comment := range comments.Comments {
			comment.PostTitle = post.Title
			comment.PostUserName = post.UserName
			comment.UserName = user.Name
		}
	}

	return posts, nil

}

func (s *PostService) UpdatePost(ctx context.Context, post *pb.PostRequest) (*pb.PostResponse, error) {
	posts, err := s.storage.Post().UpdatePost(&pb.PostRequest{
		Id:          post.Id,
		Title:       post.Title,
		Description: post.Description,
	})
	if err != nil {
		log.Println("Failed to update post info: ", err)
	}

	return posts, nil

}

func (s *PostService) DeletePost(ctx context.Context, post *pb.IdRequest) (*pb.PostResponse, error) {
	_, err := s.storage.Post().DeletePost(&pb.IdRequest{Id: post.Id})
	if err != nil {
		log.Println("Failed to delete post info: ", err)
	}

	return &pb.PostResponse{}, nil
}
