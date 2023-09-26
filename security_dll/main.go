package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

//export AES256Decrypt
func AES256DecryptAES(cipherBS64, key, iv string) string {

	// *C.char화 시키기
	// C.GoString 하기
	// C.GoByte 확인하기
	bKey := []byte(key)
	bIv := []byte(iv)

	cipherString, err := base64.StdEncoding.DecodeString(cipherBS64)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		fmt.Println("key error")
	}

	if len(cipherString) == 0 {
		fmt.Println("Empty")
	}

	mode := cipher.NewCBCDecrypter(block, bIv)
	decryptText := make([]byte, len(cipherString))
	mode.CryptBlocks(decryptText, cipherString)

	return string(UnpaddingPKCS7(decryptText, block.BlockSize()))
}

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

func main() {}

// 빌드하는 명령어
// go build -o [dll파일이름].dll -buildmode=c-shared
