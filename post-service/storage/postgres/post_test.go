package postgres

// import (
// 	"fmt"
// 	"projects/post-service/config"
// 	p "projects/post-service/genproto/post"
// 	"projects/post-service/pkg/db"
// 	"projects/post-service/storage/repo"
// 	"testing"

// 	"github.com/stretchr/testify/suite"
// )

// type PostSuiteTest struct {
// 	suite.Suite
// 	CleanUpFunc func()
// 	Repo        repo.PostStorageI
// }

// func (s *PostSuiteTest) SetupSuite() {
// 	pgPoll, cleanUp := db.ConnectDBForSuite(config.Load())
// 	s.Repo = NewPostRepo(pgPoll)
// 	s.CleanUpFunc = cleanUp
// }


// func (s *PostSuiteTest) TestPostCrud() {
// 	post := &p.PostRequest{
// 		Title: "Title for test",
// 		Description: "Description for test",
// 		UserId: "91b7870f-aeda-4b52-9328-34959fbeb09b",
// 	}

// 	createPostResp,err := s.Repo.Create(post)
// 	s.Nil(err)
// 	s.NotNil(createPostResp)
// 	s.Equal(createPostResp.Description,post.Description)
// 	s.Equal(createPostResp.Title,post.Title)
// 	s.Equal(createPostResp.UserId,post.UserId)
// 	fmt.Println(createPostResp.Id)
// 	getPostResp,err := s.Repo.GetPostById(&p.IdRequest{Id: createPostResp.Id})
// 	s.Nil(err)
// 	s.NotNil(getPostResp)
// 	s.Equal(getPostResp.Description,post.Description)
// 	s.Equal(getPostResp.Title,post.Title)
// 	s.Equal(getPostResp.UserId,post.UserId)

// 	updatePost := &p.PostRequest{
// 		Id: createPostResp.Id,
// 		Title: "New title",
// 		Description: "New description",
// 		UserId: "91b7870f-aeda-4b52-9328-34959fbeb09b",
// 	}
// 	updatePostResp, err := s.Repo.UpdatePost(updatePost)
// 	s.Nil(err)
// 	s.NotNil(updatePostResp)
// 	s.NotEqual(updatePostResp.Description,post.Description)
// 	s.NotEqual(updatePostResp.Title,post.Title)
// 	s.Equal(updatePostResp.UserId,post.UserId)

// 	listPostResp,err := s.Repo.GetPosts(&p.GetForPosts{Limit: 1,Page: 20})
// 	s.Nil(err)
// 	s.NotNil(listPostResp)

// 	_,err = s.Repo.DeletePost(&p.IdRequest{Id: createPostResp.Id})
// 	s.Nil(err)

// 	getDeletedPost,err := s.Repo.GetPostById(&p.IdRequest{Id: createPostResp.Id})
// 	s.NotNil(err)
// 	s.Nil(getDeletedPost)

// }


// func (suite *PostSuiteTest) TearDownSuite() {
// 	suite.CleanUpFunc()
// }

// func TestPostRepositoryTestSuite(t *testing.T) {
// 	suite.Run(t,new(PostSuiteTest))
// }

