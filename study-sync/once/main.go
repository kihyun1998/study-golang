package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Once : 특정 함수를 딱 한번만 실행할 때 사용함.

복잡한 for문 속에서 각종 초기화를 할 때 유용함.
*/

func hello() {
	fmt.Println("Hello")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	once := new(sync.Once)

	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Println("goroutine : ", n)
			once.Do(hello) // 고루틴 3개지만 hello는 한 번만 실행
		}(i)
	}
	fmt.Scanln()
}
