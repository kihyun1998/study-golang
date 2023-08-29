package example

import (
	"fmt"
	"time"
)

// 동기 채널2
func example2(){
	done := make(chan bool) //동기 bool 채널
	count := 3

	go func(){
		for i:=0; i<count;i++{
			done <- true				//고루틴에 true 보내고 값 꺼낼 때까지 대기함. [<-done할 때 까지 대기임]
			fmt.Println("고루틴 : ",i)
			time.Sleep(1*time.Second)
		}
	}()

	for i:=0; i<count;i++{
		<-done							//값이 들어올 때까지 대기한다. 값이 들어온다면 값을 꺼냄
		fmt.Println("메인 : ",i)
	}
}

//그래서 고(채널에 값 보냄,꺼낼 때까지 대기) > 메인(채널에서 값 꺼냄, 값 들어올 때까지 대기) > 고 > 메인 ..