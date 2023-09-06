package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Greeter struct{}

type Reply struct {
	r string
}

func (g *Greeter) Greet(name *string, reply *Reply) error {

	reply.r = fmt.Sprintf("Hello, %s!", *name)

	return nil
}

func main() {

	addr := "localhost:12345"

	greeter := new(Greeter)

	rpc.Register(greeter)

	listener, err := net.Listen("tcp", addr)

	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("RPC server listening on %s\n", addr)

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		defer conn.Close()
		go rpc.ServeConn(conn)
	}
}
