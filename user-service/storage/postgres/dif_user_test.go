package postgres

import (
	"fmt"
	"projects/user-service/storage/repo"
	"testing"

	ast "github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	asrt := ast.New(t)
	tests := []struct {
		name    string
		input   repo.UserRequest
		want    repo.UserResponse
		wantErr error
	}{
		{
			name: "error case",
			input: repo.UserRequest{
				Name:     "name for test",
				Email:    "email for test",
				Password: "password",
				UserName: "ismailjan",
			},
			want: repo.UserResponse{
				Name:     "name for test",
				Email:    "email for test",
				Password: "password",
				UserName: "ismailjan",
			},
			wantErr: invalidArgumentError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.input.Id = "invalid_uuid"
			_, err := pgRepo.Create(tc.input)
			fmt.Println(err.Error() == tc.wantErr.Error(), err)
			if err != nil && tc.wantErr != nil {
				asrt.True(tc.wantErr.Error() == err.Error())
			}

			// tc.want.Id = got.Id
			// got.CreatedAt = tc.want.CreatedAt
			// got.UpdatedAt = tc.want.UpdatedAt
			// if !reflect.DeepEqual(tc.want, got) {
			// 	t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)

			// }
		})
	}

}

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
