package biz

import (
	"context"
	pb "school/api"
)

type UserPO interface {
	GetUserByID(context.Context, int32) (*pb.User, error)
}

type User struct {
	userPO UserPO
}

func (user *User) GetUserByID(ctx context.Context, id int32) (*pb.User, error) {
	return user.userPO.GetUserByID(ctx, id)
}

func NewUser(userPO UserPO) *User {
	return &User{userPO: userPO}
}
