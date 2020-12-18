package school

import (
	"context"
	"net"
	pb "school/api"

	"google.golang.org/grpc"
)

type UserBiz interface {
	GetUserByID(context.Context, int32) (*pb.User, error)
}

type School struct {
	useBiz UserBiz
	server *grpc.Server
}

func (school *School) GetUserByID(ctx context.Context, userID *pb.UserID) (*pb.User, error) {
	return school.useBiz.GetUserByID(ctx, userID.Id)
}

// Start ...
func (school *School) Start() error {
	server := school.server
	pb.RegisterSchoolServer(server, school)
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		return err
	}

	if err := server.Serve(listen); err != nil {
		return err
	}

	return nil
}

func (school *School) Stop() {
	if school.server != nil {
		school.server.GracefulStop()
	}
}

func NewSchool(useBiz UserBiz) *School {
	return &School{useBiz: useBiz, server: grpc.NewServer()}
}
