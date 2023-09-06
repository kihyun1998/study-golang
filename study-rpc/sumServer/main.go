package main

import (
	"fmt"
	"net"
	"net/rpc"
)

/*
func Register(rcvr any) error : RPC로 사용할 함수 등록

func ServeConn(conn io.ReadWriteCloser) : 연결해서 실행

*/

// 덧셈 함수 프로시저

type Calc struct{}

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

func (c *Calc) Sum(args Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}

func main() {
	calc := new(Calc)
	rpc.Register(calc)

	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	fmt.Println("Start 6000 port")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error connection : ", err)
			continue
		}
		defer conn.Close()
		go rpc.ServeConn(conn)
	}
}
