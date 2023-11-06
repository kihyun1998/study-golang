package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	vars := [2]string{"a", "b"}

	for _, v := range vars {
		go count(v, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func count(s string, c chan bool) {
	time.Sleep(time.Second * 1)
	fmt.Println(s)
	c <- true
}
