package postgres

import (
	"database/sql"
	"fmt"
	c "projects/comment-service/genproto/comment"
	"time"
)

type CommentRepo struct {
	db *sql.DB
}

func NewCommentRepo(db *sql.DB) *CommentRepo {
	return &CommentRepo{
		db: db,
	}
}

func (s *CommentRepo) CreateComment(comment *c.CommentRequest) (*c.CommentResponse, error) {
	var res c.CommentResponse
	err := s.db.QueryRow("INSERT INTO comments(post_id,user_id,text) VALUES($1,$2,$3) RETURNING id ,post_id,user_id,text,created_at", comment.PostId, comment.UserId, comment.Text).Scan(
		&res.Id, &res.PostId, &res.UserId, &res.Text, &res.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *CommentRepo) GetComments(comment *c.ForGetComments) (*c.Comments, error) {
	offset := (comment.Page - 1) * comment.Limit
	var res c.Comments
	rows, err := s.db.Query("SELECT id, post_id, user_id, text FROM comments LIMIT $1 OFFSET $2", comment.Limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var resp c.CommentResponse
		err := rows.Scan(
			&resp.Id, &resp.PostId, &resp.UserId, &resp.Text,
		)
		if err != nil {
			return nil, err
		}
		res.Comments = append(res.Comments, &resp)
	}

	return &res, nil
}

func (s *CommentRepo) GetCommentsForPost(comment *c.GetAllCommentsRequest) (*c.Comments, error) {
	var res c.Comments
	rows, err := s.db.Query("SELECT id, post_id, user_id, text, created_at FROM comments WHERE post_id=$1", comment.PostId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var resp c.CommentResponse
		err := rows.Scan(
			&resp.Id, &resp.PostId, &resp.UserId, &resp.Text, &resp.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		res.Comments = append(res.Comments, &resp)
	}

	return &res, nil
}

func (s *CommentRepo) DeleteComment(comment *c.IdRequest) (*c.CommentResponse, error) {
	var res c.CommentResponse
	err := s.db.QueryRow("UPDATE comments SET deleted_at=$1 WHERE id=$2 RETURNING id,post_id,user_id,text,created_at", time.Now(), comment.Id).Scan(
		&res.Id, &res.PostId, &res.UserId, &res.Text, &res.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *CommentRepo) UpdateComment(comment *c.ForUpdate) (*c.CommentResponse, error) {
	var res c.CommentResponse
	err := s.db.QueryRow("UPDATE comments SET post_id=$1,text=$2,user_id=$3 WHERE id=$4 RETURNING id,post_id,post_title,user_id,user_name,post_user_name,text,created_at", comment.PostId, comment.Text, comment.UserId, comment.Id).Scan(
		&res.Id, &res.PostId, &res.PostTitle, &res.UserId, &res.UserName, &res.PostUserName, &res.Text, &res.CreatedAt,
	)
	if err != nil {
		fmt.Println("Failed to update comment: ", err)
	}

	return &res, nil
}

func (s *CommentRepo) GetComment(comment *c.GetAllCommentsRequest) (*c.Comments, error) {
	var res c.Comments
	rows, err := s.db.Query("SELECT id, post_id, user_id, text FROM comments WHERE post_id=$1", comment.PostId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var resp c.CommentResponse
		err := rows.Scan(
			&resp.Id, &resp.PostId, &resp.UserId, &resp.Text,
		)
		if err != nil {
			return nil, err
		}
		res.Comments = append(res.Comments, &resp)
	}

	return &res, nil
}
