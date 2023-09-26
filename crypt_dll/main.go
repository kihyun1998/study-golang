package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"unsafe"
)

func UnpaddingPKCS7(cipherText []byte, blockSize int) []byte {
	lastNum_byte := cipherText[len(cipherText)-1]
	lastNum := int(lastNum_byte)

	if lastNum == 0 || lastNum > len(cipherText) {
		return nil
	}

	for i := 0; i < lastNum; i++ {
		if cipherText[len(cipherText)-lastNum+i] != lastNum_byte {
			return nil
		}
	}
	return cipherText[:len(cipherText)-lastNum]
}

//export CryptFree
func CryptFree(rst *C.char) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

	fmt.Println("================CryptFree================")
	fmt.Println(rst)
	C.free(unsafe.Pointer(rst)) // Free the memory allocated by C.CString
	fmt.Println("============================================")
}

//export AesDecrypt
func AesDecrypt(cCipherBase64Text *C.char, cKey *C.char, cIV *C.char) *C.char {

	sKey := C.GoString(cKey)
	sIV := C.GoString(cIV)
	sCipherBase64Text := C.GoString(cCipherBase64Text)

	bKey := []byte(sKey)
	bIV := []byte(sIV)

	cipherString, err := base64.StdEncoding.DecodeString(sCipherBase64Text)
	if err != nil {
		fmt.Println("type error")
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		fmt.Println("key error")
	}

	if len(cipherString) == 0 {
		fmt.Println("Empty")
	}

	mode := cipher.NewCBCDecrypter(block, bIV)
	decryptText := make([]byte, len(cipherString))
	mode.CryptBlocks(decryptText, cipherString)

	rst := C.CString(string(UnpaddingPKCS7(decryptText, block.BlockSize())))

	return rst
}

func main() {}

// 빌드하는 명령어
// go build -o [dll파일이름].dll -buildmode=c-shared
