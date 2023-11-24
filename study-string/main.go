package main

import "fmt"

func main() {
	str := "real 문자열의 길이는 ?"

	runes := []rune(str)
	fmt.Printf("rune : %v", runes)
}
