package example

import "fmt"

func sum(a,b int, c chan int){
	// int형 채널에 값 보냄
	c <- a+b
}

// 동기 채널1
func example1(){
	//채널 make로 공간 할당해야함.
	c := make(chan int)

	go sum (1,2,c)

	// 채널 c에서 값을 꺼내서 n에 할당
	n:=<-c
	fmt.Println(n)
}