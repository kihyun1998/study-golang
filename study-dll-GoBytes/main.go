package main

/*
#include <stdint.h>
#include <stdlib.h>
*/

import "C"

//export ConvertToGoBytes
func ConvertToGoBytes(data *C.char) *C.char {
	return C.CString(C.GoString(data))
}

func main() {}

//Free 함수는 구현 안했음
