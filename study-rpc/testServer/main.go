package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type Reply struct {
	R string
}

type Greeter int

func (g *Greeter) Greet(name *string, reply *Reply) error {

	reply.R = fmt.Sprintf("Hello, %s!", *name)
	return nil
}

func main() {

	greeter := new(Greeter)
	rpc.Register(greeter)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":12345")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	defer listener.Close()
	fmt.Printf("RPC server listening on %s\n", ":12345")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		defer conn.Close()

		go rpc.ServeConn(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
