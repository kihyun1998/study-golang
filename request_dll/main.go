package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"unsafe"
)

func Get(uri, cookie string) (string, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Cookie", cookie)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	str := string(bytes)

	return str, nil
}

func PostPutDelete(method, uri, cookie, msg string) (string, error) {

	req, err := http.NewRequest(method, uri, bytes.NewBufferString(msg))
	if err != nil {
		return "", err
	}

	req.Header.Add("Cookie", cookie)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("connection", "close")

	//인증서 건너뛰기////
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := &http.Client{
		Transport: transport,
	}
	//////////////////

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	str := string(bytes)

	return str, nil
}

//export RequestGet
func RequestGet(cURI *C.char, cCookie *C.char) *C.char {
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("GET OPEN ERROR", r)
        }
    }()
	uri := C.GoString(cURI)
	cookie := C.GoString(cCookie)

	str, err := Get(uri, cookie)
	if err != nil {
		rst := C.CString("[Error] : " + err.Error())
		return rst
	}
	rst := C.CString(string([]rune(str)))
	fmt.Println("------------------create---------------------------------")
	fmt.Println(rst)
	fmt.Println("-------------------------------------------------------")
	
	return rst
}

//export RequestPPD
func RequestPPD(cMethod *C.char, cURI *C.char, cCookie *C.char, cMsg *C.char) *C.char {

	method := C.GoString(cMethod)
	uri := C.GoString(cURI)
	cookie := C.GoString(cCookie)
	msg := C.GoString(cMsg)

	str, err := PostPutDelete(method, uri, cookie, msg)
	if err != nil {
		rst := C.CString("[Error] : " + err.Error())
		return rst
	}
	rst:= C.CString(string([]rune(str)))

	return rst
}

//export Free
func Free(rst *C.char) {
	// 주소값을 C#에서 받아서 그걸 Free해주도록
	// 전역변수 쓰지마
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("OPEN ERROR", r)
        }
    }()
	fmt.Println("------------------free---------------------------------")
	fmt.Println(rst)
	C.free(unsafe.Pointer(rst))
	fmt.Println("-------------------------------------------------------")
}

func main() {}
