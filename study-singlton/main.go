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

func ShowItem() {
	instance := NewMap()
	for k, v := range instance {
		fmt.Println("key : ", k, "value : ", v)
	}
}

func main() {
	AddValueByKey("a", "b")
	AddValueByKey("c", "d")

	ShowItem()
}
