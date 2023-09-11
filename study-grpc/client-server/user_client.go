package client_server

import (
	user_proto "study-golang/study-grpc/protos/v1/user"
	"sync"

	"google.golang.org/grpc"
)

var (
	once   sync.Once
	client user_proto.UserClient
)

func GetUserClient(serviceHost string) user_proto.UserClient {
	once.Do(func() {
		conn, _ := grpc.Dial(serviceHost,
			grpc.WithInsecure(), // transport security 비활성화
			grpc.WithBlock(),
		)

		client = user_proto.NewUserClient(conn)
	})
	return client
}
