package data

import (
	"database/sql"

	pb "school/api"

	"context"
)

type User struct {
	db *sql.DB
}

func (user *User) GetUserByID(ctx context.Context, id int32) (*pb.User, error) {
	return &pb.User{Id: 22, Name: "hello"}, nil
}

func NewUser() *User {
	return &User{}
}
