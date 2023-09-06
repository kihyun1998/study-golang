package main

import (
	"fmt"
	"net/rpc"
)

/*
func Dial(network, address string) (*Client, error) : 프로토콜, IP 주소, 포트 번호를 설정하여 RPC 서버에 연결

func (client *Client) Call(serviceMethod string, args any, reply any) error : RPC 서버의 함수를 호출(동기)

func (client *Client) Go(serviceMethod string, args any, reply any, done chan *Call) *Call : RPC 서버의 함수를 고루틴으로 호출(비동기)
*/

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

func main() {

	fmt.Println("===시작===")
	client, err := rpc.Dial("tcp", "127.0.0.1:6000")
	fmt.Println(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer client.Close()

	args := &Args{1, 2}
	reply := new(Reply)

	err = client.Call("Calc.Sum", args, reply)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("sync : ", reply.C)

	args.A = 4
	args.B = 9
	sumCall := client.Go("Calc.Sum", args, reply, nil)

	fmt.Println(sumCall)
	<-sumCall.Done
	fmt.Println(sumCall)
	fmt.Println("async : ", reply.C)
}
