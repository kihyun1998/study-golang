package main

import (
	"context"
	"log"
	"study-golang/study-grpc/data"
	post_proto "study-golang/study-grpc/protos/v1/post"
	user_proto "study-golang/study-grpc/protos/v1/user"

	"google.golang.org/grpc"
)

const port = "5000"

type postServer struct {
	post_proto.PostServer
	userClient user_proto.UserClient
}

func (s *postServer) ListPostsByUserId(ctx context.Context, req *post_proto.ListPostsByUserIdRequest) (*post_proto.ListPostsByUserIdResponse, error) {
	userId := req.UserId

	res, err := s.userClient.GetUser(ctx, &user_proto.GetUserRequest{UserId: userId})
	if err != nil {
		return nil, err
	}

	var postMessages []*post_proto.PostMessage

	for _, uPost := range data.UserPosts {
		if uPost.UserId != userId {
			continue
		}

		for _, post := range uPost.Posts {
			post.Author = res.UserMessage.Name
		}

		postMessages = uPost.Posts
		break
	}

	return &post_proto.ListPostsByUserIdResponse{
		PostMessages: postMessages,
	}, nil
}

func (s *postServer) ListPosts(ctx context.Context, req *post_proto.ListPostsByUserIdRequest)(res *post_proto.ListPostsByUserIdResponse, error){
	var postMessages = []*post_proto.PostMessage

	for _,uPost := range data.UserPosts{
		res,err := s.userClient.GetUser(ctx,&user_proto.GetUserRequest{UserId: uPost.UserId})
		if err != nil{
			return nil,err
		}

		for _,post := range uPost.Posts{
			p.Author = res.UserMessage.Name
		}
	}
}

func main() {
	listen, err := net.Listen("tcp",":"+port)
	if err!=nil{
		log.Fatalf("Failed : %v",err)
	}

	isUserClient := client.GetUserClient("localhost:5000")
	grpcServer := grpc.NewServer()
	post_proto.RegisterPostServer(grpcServer,&postServer{
		userClient: isUserClient,
	})
}
