package main

//client.go

import (
	"log"

	pb "school/api"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSchoolClient(conn)

	r, err := c.GetUserByID(context.Background(), &pb.UserID{Id: 22})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Printf("user: %#v", r)
}
