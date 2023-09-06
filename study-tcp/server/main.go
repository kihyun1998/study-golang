package main

/*
func Listen(net, addr string)(Listener, error) : 프로토콜, IP주소, 포트번호를 설정하여 네트워크 연결을 대기
- net 종류 : tcp, tcp4, tcp6, unix, unixpacket

func (l *TCPListener) Accept() (Conn, error) : 클라이언트가 연결되면 TCP 연결을 리턴 (Conn == 커넥션)

func (l *TCPListener) Close() error : TCP 연결 대기를 닫음

func (c *TCPConn) Read(b []byte) (int, error) : 받은 데이터를 읽음

func (c *TCPConn) Write(b []byte) (int, error) : 데이터를 보냄

func (c *TCPConn) Close() error : TCP 연결을 닫음

Conn == 연결된거
TCPListener == 연결 대기중인거

*/

import (
	"fmt"
	"net"
)

func requestHandler(c net.Conn) {
	data := make([]byte, 4096)

	for {
		n, err := c.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(n)
		fmt.Println("response data : " + string(data[:n]))

		_, err = c.Write(data[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	//ln값도 메모리 주소로 생김 like &{0xc00010ca00 { <nil> 0}}
	// 위의 예시 값 중 두번째 값은 모르겠음
	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("ln is")
	fmt.Println(ln)

	defer ln.Close()

	for {
		// 실제 conn값은 &{{0xc001111}}같이 메모리 주소가 할당된다.
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()

		fmt.Print("conn is ")
		fmt.Println(conn)
		go requestHandler(conn)
	}
}
