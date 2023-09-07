package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type Quotient struct {
	Quo, Rem string
}

type Arith int

func (t *Arith) Divide(name *string, quo *Quotient) error {
	quo.Quo = fmt.Sprintf("Hello, %s !\n", *name)
	quo.Rem = fmt.Sprintf("I am Good %s\n", *name)
	return nil
}

func main() {

	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
