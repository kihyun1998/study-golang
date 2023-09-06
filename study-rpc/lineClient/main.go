package main

import (
	"bufio"
	"log"
	"net/rpc"
	"os"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:5000")
	if err != nil {
		log.Fatal(err)
	}

	in := bufio.NewReader(os.Stdin) // 입력
	for {
		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		var reply bool
		err = client.Call("Listener.GetLine", line, &reply)
		if err != nil {
			log.Fatal(err)
		}
	}
}
