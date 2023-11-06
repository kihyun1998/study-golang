package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	vars := [4]string{"a", "b", "c", "d"}

	for _, v := range vars {
		go count(v, c)
	}
	for i := 0; i < len(vars); i++ {
		fmt.Println(<-c)
	}
}

func count(s string, c chan string) {
	time.Sleep(time.Second * 1)
	c <- s + " here"
}
