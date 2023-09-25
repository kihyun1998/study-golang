package data

import (
	user_proto "study-grpc/proto/v1/user"
)

var UserData = []*user_proto.UserMessage{
	{
		UserId:      "1",
		Name:        "Thomas",
		PhoneNumber: "010-1111-1111",
		Age:         22,
	},
	{
		UserId:      "2",
		Name:        "Mark",
		PhoneNumber: "010-2222-2222",
		Age:         55,
	},
	{
		UserId:      "3",
		Name:        "John",
		PhoneNumber: "010-3333-3333",
		Age:         15,
	},
	{
		UserId:      "4",
		Name:        "Leo",
		PhoneNumber: "010-4444-4444",
		Age:         37,
	},
	{
		UserId:      "5",
		Name:        "Kim",
		PhoneNumber: "010-5555-5555",
		Age:         25,
	},
}
