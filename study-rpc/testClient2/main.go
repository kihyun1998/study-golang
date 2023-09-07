package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem string
}

func main() {

	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	name := "park"
	var quot Quotient
	err = client.Call("Arith.Divide", name, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("%s%s", quot.Quo, quot.Rem)

}
