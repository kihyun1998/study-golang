package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var data int = 0
	var rwMutex = new(sync.RWMutex)

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock() // 쓰기 mutex lock

			data += 1
			fmt.Println("write : ", data)
			time.Sleep(10 * time.Millisecond)

			rwMutex.Unlock() // 쓰기 mutex unlock
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock() // 읽기 mutex lock

			fmt.Println("read 1 : ", data)
			time.Sleep(1 * time.Second)

			rwMutex.RUnlock() // 읽기 mutex unlock
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock() // 읽기 mutex lock

			fmt.Println("read 2 : ", data)
			time.Sleep(2 * time.Second)

			rwMutex.RUnlock() // 읽기 mutex unlock
		}
	}()

	time.Sleep(10 * time.Second)
}
