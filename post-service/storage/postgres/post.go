package postgres

import (
	"database/sql"
	"fmt"
	p "projects/post-service/genproto/post"
)

type postRepo struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *postRepo {
	return &postRepo{db: db}
}

func (r *postRepo) Create(post *p.PostRequest) (*p.PostResponse, error) {
	var resp p.PostResponse
	err := r.db.QueryRow(`
		INSERT INTO 
			posts(title, description, user_id) VALUES($1,$2,$3) 
		RETURNING 
			id,
			title,
			description,
			user_id,
			created_at,
			updated_at
		`, post.Title, post.Description, post.UserId).Scan(
		&resp.Id, &resp.Title, &resp.Description, &resp.UserId, &resp.CreatedAt, &resp.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *postRepo) GetPostById(post *p.IdRequest) (*p.PostResponse, error) {
	var postResp p.PostResponse
	err := r.db.QueryRow(`
	SELECT id,title,description,user_id,created_at,updated_at from posts  WHERE user_id=$1`, post.Id).Scan(
		&postResp.Id, &postResp.Title, &postResp.Description, &postResp.UserId, &postResp.CreatedAt, &postResp.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &postResp, nil
}

func (r *postRepo) GetAllPostsByUserId(post *p.IdRequest) (*p.Posts, error) {
	var postsResp p.Posts
	rows, err := r.db.Query("SELECT id,title,description,user_id,created_at,updated_at FROM posts WHERE user_id=$1", post.Id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var postResp p.PostResponse
		err := rows.Scan(
			&postResp.Id,
			&postResp.Title,
			&postResp.Description,
			&postResp.UserId,
			&postResp.CreatedAt,
			&postResp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		postsResp.Posts = append(postsResp.Posts, &postResp)
	}

	return &postsResp, nil

}

func (r *postRepo) GetPostForComment(id *p.IdRequest) (*p.PostResponse, error) {
	var postResp p.PostResponse
	err := r.db.QueryRow(`
	SELECT id,title,description,user_id,created_at,updated_at from POSTS WHERE id=$1`, id.Id).Scan(
		&postResp.Id, &postResp.Title, &postResp.Description, &postResp.UserId, &postResp.CreatedAt, &postResp.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &postResp, nil
}

func (r *postRepo) GetPostForUser(id *p.IdRequest) (*p.Posts, error) {
	var postsResp p.Posts
	rows, err := r.db.Query("SELECT id,title,description,user_id,created_at,updated_at FROM posts WHERE user_id=$1", id.Id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var postResp p.PostResponse
		err := rows.Scan(
			&postResp.Id,
			&postResp.Title,
			&postResp.Description,
			&postResp.UserId,
			&postResp.CreatedAt,
			&postResp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		postsResp.Posts = append(postsResp.Posts, &postResp)
	}

	return &postsResp, nil

}

func (r *postRepo) GetPosts(post *p.GetForPosts) (*p.Posts, error) {
	var resp p.Posts
	offset := (post.Page - 1) * post.Limit

	rows, err := r.db.Query("SELECT id,title,description,user_id,created_at,updated_at FROM posts LIMIT $1 OFFSET $2", post.Limit, offset)
	if err != nil {
		fmt.Println("Failed to select all posts with limit and page", err)
	}

	for rows.Next() {
		var res p.PostResponse
		err := rows.Scan(
			&res.Id, &res.Title, &res.Description, &res.UserId, &res.CreatedAt, &res.UpdatedAt,
		)
		if err != nil {
			fmt.Println("Failed to scan post user info with limit and page", err)
		}
		resp.Posts = append(resp.Posts, &res)
	}

	return &resp, nil
}

func (r *postRepo) UpdatePost(post *p.PostRequest) (*p.PostResponse, error) {
	var res p.PostResponse
	fmt.Println(post.Id)
	err := r.db.QueryRow("UPDATE posts SET title=$1, description=$2 WHERE id=$3 RETURNING title,description,user_id,id", post.Title, post.Description, post.Id).Scan(
		&res.Title, &res.Description,&res.UserId,&res.Id,
	)
	if err != nil {
		fmt.Println("Failed to update post info: ", err)
	}



	return &res, nil
}

func (r *postRepo) DeletePost(post *p.IdRequest) (*p.PostResponse, error) {
	err := r.db.QueryRow("DELETE FROM posts WHERE id=$1", post.Id)
	if err != nil {
		fmt.Println("Failed to delete post info: ", err)
	}

	return &p.PostResponse{}, nil
}
