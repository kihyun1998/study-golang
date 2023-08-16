package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func hello(n int) {
// 	r := rand.Intn(100)
// 	time.Sleep(time.Duration(r))
// 	fmt.Println(n)
// }

// func main() {
// 	for i := 0; i < 100; i++ {
// 		go hello(i)
// 	}
// 	fmt.Scanln()
// }

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(0))

	s := "hello"

	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i) // << 얘가 함수 인자값
	}
	fmt.Scanln()
}
