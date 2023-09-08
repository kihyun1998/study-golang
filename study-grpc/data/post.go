package data

import (
	post_proto "study-golang/study-grpc/protos/v1/post"
)

type PostData struct {
	UserId string
	Posts  []*post_proto.PostMessage
}

var UserPosts = []*PostData{
	{
		UserId: "1",
		Posts: []*post_proto.PostMessage{
			{
				PostId: "1",
				Author: "",
				Title:  "안녕하세요",
				Body:   "자기소개 시작 자기소개 끝",
				Tags:   []string{"hello", "world"},
			},
			{
				PostId: "2",
				Author: "",
				Title:  "~~에 대하여",
				Body:   "~~는 물결표시고 1 옆에 있습니다.",
				Tags:   []string{"1", "~", "물결"},
			},
		},
	},
}
