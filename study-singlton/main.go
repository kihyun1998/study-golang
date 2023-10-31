package main

import (
	"fmt"
	"sync"
)

// once는 한번만 실행하게 해준다.
var once sync.Once

type singleton map[string]string

var instance singleton

func NewMap() singleton {
	// 한번만 실행
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}

func AddValueByKey(key string, value string) {
	instance := NewMap()
	instance[key] = value
}

func main() {
	AddValueByKey("a", "b")
	AddValueByKey("c", "d")

	globalMap := NewMap()
	fmt.Println("키 'a'의 값:", globalMap["a"])
	fmt.Println("키 'c'의 값:", globalMap["c"])
}
