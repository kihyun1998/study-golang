package main

import (
	"log"
	"net"

	data "study-golang/study-grpc/data"
	user_proto "study-golang/study-grpc/protos/v1/user"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = "5000"

type userServer struct {
	user_proto.UserServer
}

func (s *userServer) GetUser(ctx context.Context, req *user_proto.GetUserRequest) (*user_proto.GetUserResponse, error) {
	userId := req.UserId                    // 요청한 user Id
	var userMessage *user_proto.UserMessage // 유저 정보를 담을 변수
	for _, u := range data.UserData {       // 요청한 userId를 찾기 위한 탐색
		if u.UserId != userId {
			continue
		}
		userMessage = u // 찾으면 유저 정보 담기
		break
	}
	return &user_proto.GetUserResponse{ // 응답으로 유저 정보, 에러는 nil
		UserMessage: userMessage,
	}, nil
}

func (s *userServer) ListUsers(ctx context.Context, req *user_proto.ListUsersRequest) (*user_proto.ListUsersResponse, error) {
	userMessage := make([]*user_proto.UserMessage, len(data.UserData))

	// for i, u := range data.UserData {
	// 	userMessage[i] = u
	// }
	copy(userMessage, data.UserData) // 위에 for문과 같은 동작을 한다.

	return &user_proto.ListUsersResponse{
		UserMessages: userMessage,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	user_proto.RegisterUserServer(grpcServer, &userServer{})

	log.Printf("Start gRPC server on %s port", port)

	// if문을 이런식으로 사용하면 err변수는 이 if문 안에서만 사용되고 없어짐
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %s ", err)
	}
}
