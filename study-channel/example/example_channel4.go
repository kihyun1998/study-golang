package example

import (
	"fmt"
)

func example4(){
	c := make(chan int)
	go func(){
		for i:=0; i<5; i++{
			c <- i //채널 값 보냄
		}
		close(c) //닫기
	}()

	for i:= range c{
		fmt.Println(i)
	}
}
