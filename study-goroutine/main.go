package main

import (
	"fmt"
	"runtime"
)

// 클로저를 go 루틴으로 실행
func main() {
	runtime.GOMAXPROCS(1)

	s := "hello"

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(s, i)

		}()
	}
	fmt.Scanln()
}
