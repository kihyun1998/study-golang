package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
	"unsafe"
)

var Rst *C.char

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

	uri := C.GoString(cURI)
	cookie := C.GoString(cCookie)

	str, err := Get(uri, cookie)
	if err != nil {
		Rst := C.CString("[Error] : " + err.Error())
		return Rst
	}
	Rst := C.CString(string([]rune(str)))

	return Rst
}

//export RequestPPD
func RequestPPD(cMethod *C.char, cURI *C.char, cCookie *C.char, cMsg *C.char) *C.char {

	method := C.GoString(cMethod)
	uri := C.GoString(cURI)
	cookie := C.GoString(cCookie)
	msg := C.GoString(cMsg)

	str, err := PostPutDelete(method, uri, cookie, msg)
	if err != nil {
		Rst := C.CString("[Error] : " + err.Error())
		return Rst
	}
	Rst := C.CString(string([]rune(str)))

	return Rst
}

//export Free
func Free() {
	C.free(unsafe.Pointer(Rst))
}

func main() {}
