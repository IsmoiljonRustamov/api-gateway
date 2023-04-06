package postgres

// pb "projects/user-service/genproto/user"

// type UserSuiteTest struct {
// 	suite.Suite
// 	CleanUpFunc func()
// 	Repo        repo.UserStorageI
// }

// func (s *UserSuiteTest) SetupSuite() {
// 	pgPoll, cleanUp := db.ConnectDBForSuite(config.Load())
// 	s.Repo = NewUserRepo(pgPoll)
// 	s.CleanUpFunc = cleanUp
// }

// func (s *UserSuiteTest) TestUserCrud() {
// 	user := &pb.UserRequest{
// 		Name: "Ismoiljon",
// 		Email: "ismoiljonrustamov6@gmail.com",
// 		Password: "12",
// 		UserName: "ismailjan",
// 	}
// 	createUserResp,err := s.Repo.Create(user)
// 	s.Nil(err)
// 	s.NotNil(createUserResp)
// 	s.Equal(user.Name,createUserResp.Name)
// 	s.Equal(user.Email,createUserResp.Email)
// 	s.Equal(user.Password,createUserResp.Password)
// 	s.Equal(user.UserName,createUserResp.UserName)

// 	getUserResp,err := s.Repo.GetUserById(&pb.IdRequest{Id: createUserResp.Id})
// 	s.Nil(err)
// 	s.NotNil(getUserResp)
// 	s.Equal(getUserResp.Name,user.Name)
// 	s.Equal(getUserResp.Email,user.Email)
// 	s.Equal(getUserResp.Password, user.Password)
// 	s.Equal(getUserResp.UserName,user.UserName)

// 	updateUser := &pb.UserRequest{
// 		Id: createUserResp.Id,
// 		Name: "ismoiljon",
// 		Email: "ismoiljonrustamo@gmail.com",
// 		Password: "13",
// 		UserName: "Ismailjan",
// 	}
// 	updateResp,err := s.Repo.UpdateUser(updateUser)
// 	s.Nil(err)
// 	s.NotNil(updateResp)
// 	s.NotEqual(updateResp.Email,user.Email)
// 	s.NotEqual(updateResp.Name,user.Name)

// 	listUserResp,err := s.Repo.GetUsers(&pb.UserForGetUsers{Limit: 1,Page: 100})
// 	s.Nil(err)
// 	s.NotNil(listUserResp)

// 	_, err = s.Repo.DeleteUser(&pb.IdRequest{Id: createUserResp.Id})
// 	s.Nil(err)

// }

// func (suite *UserSuiteTest) TearDownSuite() {
// 	suite.CleanUpFunc()
// }

// // In order for 'go test ' to run this suite, we need to create
// // a normal test function and pass our 	suite to suite.Run
// func TestUserRepositoryTestSuite(t *testing.T) {
// 	suite.Run(t, new(UserSuiteTest))
// }
