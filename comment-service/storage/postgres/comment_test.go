package postgres

import (
	"projects/comment-service/config"
	p "projects/comment-service/genproto/comment"
	"projects/comment-service/pkg/db"
	"projects/comment-service/storage/repo"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CommentSuiteTest struct {
	suite.Suite
	CleanUpFunc func()
	Repo        repo.CommentStorageI
}

func (s *CommentSuiteTest) SetupSuite() {
	pgPoll, cleanUp := db.ConnectDBForSuite(config.Load())
	s.Repo = NewCommentRepo(pgPoll)
	s.CleanUpFunc = cleanUp
}

func (s *CommentSuiteTest) TestCommentCrud() {
	comment := &p.CommentRequest{
		UserId: 1,
		PostId: 2,
		Text:   "Text for test",
	}

	createCommentResp, err := s.Repo.CreateComment(comment)
	s.Nil(err)
	s.NotNil(createCommentResp)
	s.Equal(createCommentResp.UserId, comment.UserId)
	s.Equal(createCommentResp.PostId, comment.PostId)
	s.Equal(createCommentResp.Text, comment.Text)

	getCommentResp, err := s.Repo.GetComment(&p.GetAllCommentsRequest{PostId: createCommentResp.PostId})
	s.Nil(err)
	s.NotNil(getCommentResp)
	for _, comm := range getCommentResp.Comments {
		s.Equal(comm.PostId, createCommentResp.PostId)
	}

	listCommentResp, err := s.Repo.GetComments(&p.ForGetComments{Limit: 1, Page: 20})
	s.Nil(err)
	s.NotNil(listCommentResp)
	for _, comm := range listCommentResp.Comments {
		s.Equal(comm.PostId, createCommentResp.PostId)
		s.Equal(comm.UserId, createCommentResp.UserId)
		s.Equal(comm.Text, createCommentResp.Text)

	}

	_, err = s.Repo.DeleteComment(&p.IdRequest{Id: createCommentResp.Id})
	s.Nil(err)
}

func (suite *CommentSuiteTest) TearDownSuite() {
	suite.CleanUpFunc()
}

func TestPostRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CommentSuiteTest))
}
