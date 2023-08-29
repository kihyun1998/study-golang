package example

import (
	"fmt"
)

// 비동기 채널
func example3(){
	
	// bool 채널인데 버퍼가 2개인 비동기 채널
	// 버퍼를 한 개 이상 설정하면 비동기 채널
	done := make(chan bool) 
	count := 10

	go func(){
		for i:=0; i<count;i++{
			done<-true
			fmt.Println("고루틴 : ", i)
		}
	}()

	for i:=0;i<count;i++{
		<-done
		fmt.Println("메인 함수 : ",i)
	}
}