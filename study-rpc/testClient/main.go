package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Reply struct {
	R string
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	var reply Reply
	name := "park"
	err = client.Call("Greeter.Greet", name, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.R)
}
