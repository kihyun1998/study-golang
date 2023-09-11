package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Cond : 조건 변수(condition variable)이다. 대기하고 있는 하나의 객체를 깨울 수도 있고 여러 개를 동시에 깨울 수도 있습니다.

func (c *Cond) Wait() : 고루틴 실행 멈추고 대기
func (c *Cond) Signal() : 대기하고 이쓴 고루틴 하나만 깨움
func (c *Cond) Broadcast() : 대기하고 있는 모든 고루틴을 깨움
*/

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var cond = sync.NewCond(mutex)

	c := make(chan bool, 3)

	for i := 0; i < 3; i++ {
		go func(n int) {
			mutex.Lock()

			c <- true
			fmt.Println("[wait] begin : ", n)
			cond.Wait()
			fmt.Println("[wait] end : ", n)

			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c //채널에서 값 꺼내기, (고루틴 기다리기)
	}

	for i := 0; i < 3; i++ {
		mutex.Lock()
		fmt.Println("signal : ", i)
		cond.Signal()
		mutex.Unlock()
	}
	fmt.Scanln()
}
