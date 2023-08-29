package example

import (
	"fmt"
)

func example5(){
	c := make(chan int)
	go func(){
		c<-1
	}()

	// 2번째 매게변수로 채널 닫혔는지 열렸는지 확인 가능
	// 열렸으면 true 닫혔으면 false
	a, ok := <-c
	fmt.Println(a, ok)

	close(c)

	a, ok = <-c
	fmt.Println(a, ok)
}
