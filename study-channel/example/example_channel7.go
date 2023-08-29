package example

import (
	"fmt"
)

// 함수 return을 채널로

func sum7(a, b int) <-chan int{
	out:=make(chan int)
	go func(){
		out <- a+b
	}()
	return out
}

func example7(){
	c:=sum7(1,2)
	fmt.Println(<-c)
}
