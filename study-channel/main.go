package main

import (
	"fmt"
	"time"
)

// 여러 채널 제어

func main(){
	c1 := make(chan int)
	c2 := make(chan string)

	go func(){ //무한루프
		for{
			c1<-10
			time.Sleep(100*time.Millisecond)
		}
	}()

	go func(){ //무한루프
		for{
			c2 <- "Hello"
			time.Sleep(500*time.Millisecond)
		}
	}()

	go func(){ //무한루프
		for{
			select{
			case i := <-c1:
				fmt.Println("c1:",i)
			case s := <-c2:
				fmt.Println("c2:",s)
			}
		}
	}()

	time.Sleep(10*time.Second) // 10초동안 프로그램 실행
}