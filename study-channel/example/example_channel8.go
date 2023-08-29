package example

import (
	"fmt"
)

// 채널을 받아서 덧셈하고 채널을 반환

func num(a,b int) <-chan int{
	out := make(chan int)
	go func(){
		out <- a
		out <- b
		close(out)
	}()
	return out
}

func sum8(c<-chan int)<-chan int{
	out:=make(chan int)
	go func(){
		r:=0
		for i:=range c{
			r += i
		}
		out <- r
	}()
	return out
}

func example8(){
	c:=num(1,2)
	out:=sum8(c)
	fmt.Println(<-out)
}
