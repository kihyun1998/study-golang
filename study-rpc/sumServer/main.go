package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
)

/*
func Register(rcvr any) error : RPC로 사용할 함수 등록

func ServeConn(conn io.ReadWriteCloser) : 연결해서 실행

*/

// 덧셈 함수 프로시저

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

type Calc int

func (c *Calc) Sum(args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}

func main() {
	calc := new(Calc)
	rpc.Register(calc)

	// tcpAddr, err := net.ResolveTCPAddr("tcp", ":12345")
	// checkError(err)

	// listener, err := net.ListenTCP("tcp", tcpAddr)
	// checkError(err)

	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Start 12345 port")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error connection : ", err)
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
