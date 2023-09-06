package main

/*
func Dial(network, address string) (Conn, error) : 프로토콜, IP주소, 포트 번호를 설정하여 서버에 연결
network 종류 : tcp, tcp4(IPv4-only), tcp6(IPv6-only), udp, udp4(IPv4-only), udp6(IPv6-only), ip, ip4(IPv4-only), ip6(IPv6-only), unix, unixgram, unixpacket.

func (c *TCPConn) Read(b []byte) (int, error) : 받은 데이터를 읽음

func (c *TCPConn) Write(b []byte) (int, error) : 데이터를 보냄

func (c *TCPConn) Close() error : TCP 연결을 닫음
*/

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	client, err := net.Dial("tcp", "127.0.0.1:5000")
	fmt.Print("Client is ")
	fmt.Println(client)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer client.Close()

	// c == client
	go func(c net.Conn) {
		data := make([]byte, 4096)

		for {
			n, err := c.Read(data)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(string(data[:n]))

			time.Sleep(1 * time.Second)
		}
	}(client)

	go func(c net.Conn) {
		i := 0
		for {
			message := "Hello" + strconv.Itoa(i)
			_, err := c.Write([]byte(message))
			if err != nil {
				fmt.Println(err)
				return
			}
			i++
			time.Sleep(1 * time.Second)
		}
	}(client)
	fmt.Scanln()
}
