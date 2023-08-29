package example

import (
	"fmt"
)

// 보내기 전용과 받기 전용 채널

//보내기 전용 채널(send-only)
func producer(c chan<- int){
	for i:=0;i<5;i++{
		c<-i
	}
	
	c<-100
}

//받기 전용 채널(receive-only)
func consumer(c<-chan int){
	for i:= range c{
		fmt.Println(i)
	}

	fmt.Println(<-c)
}

func example6(){
	c:=make(chan int)
	go producer(c)
	go consumer(c)

	fmt.Scanln()
}
