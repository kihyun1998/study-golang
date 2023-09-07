package main

import (
	"log"
	"net"

	user_proto "study-golang/study-grpc/protos/v1/user"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const port = "5000"

type userServer struct {
	user_proto.UserServer
}

func (s *userServer) GetUser(ctx context.Context, req *user_proto.GetUserRequest) (*user_proto.GetUserResponse, error) {
	userId := req.UserId

}

func main() {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	grpcServer := grpc.NewServer()

	log.Printf("Start gRPC server on %s port", port)

	// if문을 이런식으로 사용하면 err변수는 이 if문 안에서만 사용되고 없어짐
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %s ", err)
	}
}
