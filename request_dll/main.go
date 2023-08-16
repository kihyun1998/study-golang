package main

import "C"

import (
	"io"
	"net/http"
)

type ReturnValue struct {
	Result *C.char
	Error  *C.char
}

func GETRequest(uri, cookie string) (string, error) {
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

//export GET
func GET(cURI *C.char, cCookie *C.char) *C.char {

	uri := C.GoString(cURI)
	cookie := C.GoString(cCookie)

	str, err := GETRequest(uri, cookie)
	if err != nil {
		return C.CString(err.Error())
	}
	return C.CString(string([]rune(str)))
}

func main() {}
