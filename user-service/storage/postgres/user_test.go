package postgres

// import (
// 	pb "projects/user-service/genproto/user"
// 	"reflect"
// 	"testing"
// )

// func TestUserCreate(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		input   pb.UserRequest
// 		want    pb.UserResponse
// 		wantErr bool
// 	}{
// 		{
// 			name: "succes",
// 			input: pb.UserRequest{
// 				Name:     "name for test",
// 				Email:    "email for test",
// 				Password: "password",
// 				UserName: "ismailjan",
// 			},
// 			want: pb.UserResponse{
// 				Name:     "name for test",
// 				Email:    "email for test",
// 				Password: "password",
// 				UserName: "ismailjan",
// 			},
// 			wantErr: false,
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			got, err := pgRepo.Create(&tc.input)
// 			if err != nil {
// 				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
// 			}

// 			tc.want.Id = got.Id
// 			got.CreatedAt = tc.want.CreatedAt
// 			got.UpdatedAt = tc.want.UpdatedAt
// 			if !reflect.DeepEqual(tc.want, *got) {
// 				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)

// 			}
// 		})
// 	}

// }

// func TestGetUserById(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		input   pb.IdRequest
// 		want    pb.UserResponse
// 		wantErr bool
// 	}{
// 		{
// 			name: "succes",
// 			input: pb.IdRequest{
// 				Id: 93,
// 			},
// 			want: pb.UserResponse{
// 				Id:       93,
// 				Name:     "name for test",
// 				Email:    "email for test",
// 				Password: "password",
// 				UserName: "ismailjan",
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			got, err := pgRepo.GetUserById(&tc.input)
// 			if err != nil {
// 				t.Fatalf("%s: expexted: %v. got: %v", tc.name, tc.wantErr, err)
// 			}
// 			tc.want.Id = got.Id
// 			tc.want.CreatedAt = got.CreatedAt
// 			tc.want.UpdatedAt = got.UpdatedAt
// 			if !reflect.DeepEqual(tc.want, *got) {
// 				t.Fatalf("%s: expexted: %v, got: %v", tc.name, tc.want, got)
// 			}
// 		})
// 	}
// }

// func TestUserUpdate(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		input   pb.UserRequest
// 		want    pb.UserForUpdate
// 		wantErr bool
// 	}{
// 		{
// 			name: "succes",
// 			input: pb.UserRequest{
// 				Id:       90,
// 				Name:     "new",
// 				Email:    "new",
// 				Password: "new",
// 				UserName: "newuser",
// 			},
// 			want: pb.UserForUpdate{
// 				Name:     "new",
// 				Email:    "new",
// 				Password: "new",
// 				UserName: "newuser",
// 				Id:       90,
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			got, err := pgRepo.UpdateUser(&tc.input)
// 			if err != nil {
// 				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
// 			}
// 			tc.want.Id = got.Id
// 			got.CreatedAt = tc.want.CreatedAt
// 			got.UpdatedAt = tc.want.UpdatedAt
// 			if !reflect.DeepEqual(tc.want, *got) {
// 				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
// 			}
// 		})
// 	}

// }

// func TestDeleteUser(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		input   pb.IdRequest
// 		want    pb.UserForUpdate
// 		wantErr bool
// 	}{
// 		{
// 			name: "succes",
// 			input: pb.IdRequest{
// 				Id: 106,
// 			},
// 			want: pb.UserForUpdate{
// 				Id:       106,
// 				Name:     "name for test",
// 				Email:    "email for test",
// 				Password: "password",
// 				UserName: "ismailjan",
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			got, err := pgRepo.DeleteUser(&tc.input)
// 			if err != nil {
// 				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
// 			}
// 			tc.want.CreatedAt = got.CreatedAt
// 			tc.want.UpdatedAt = got.UpdatedAt
// 			if !reflect.DeepEqual(tc.want, *got) {
// 				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)

// 			}
// 		})
// 	}
// }
