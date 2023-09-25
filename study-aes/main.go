package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"fmt"
)

func AES256Encrypt(plainString string, key string, iv string) string {

	blockSize := 256
	bKey := []byte(key)
	bIv := []byte(iv)

	paddingByte := PaddingPKCS7([]byte(plainString), blockSize)

	block, err := aes.NewCipher(bKey)
	if err != nil {
		fmt.Println("ERR")
	}

	if plainString == "" {
		fmt.Println("Empty")
	}

	cipherByte := make([]byte, len(paddingByte))
	mode := cipher.NewCBCEncrypter(block, bIv)
	mode.CryptBlocks(cipherByte, paddingByte)

	// fmt.Println("Cipher Text : ", base64.StdEncoding.EncodeToString(cipherByte))
	return base64.StdEncoding.EncodeToString(cipherByte)
}

func AES256Decode(cipherBS64, key, iv string) string {
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

	// fmt.Println("Decrypt Text : ", string(UnpaddingPKCS7(decryptText, block.BlockSize())))
	return string(UnpaddingPKCS7(decryptText, block.BlockSize()))
}

func PaddingPKCS7(plainText []byte, blockSize int) []byte {
	paddingBuf := &bytes.Buffer{}
	padLen := blockSize - (len(plainText) % blockSize)

	// binary.Write(버퍼에, LittleEndian으로, 이 []byte 값 append)
	binary.Write(paddingBuf, binary.LittleEndian, plainText)
	// bytes.Repeat(어떤 값을 반복할지 []byte, 반복할 횟수 int)
	padVar := []byte{byte(padLen)}
	binary.Write(paddingBuf, binary.LittleEndian, bytes.Repeat(padVar, padLen))

	// fmt.Println("Padding Data : ", paddingBuf.Bytes())

	return paddingBuf.Bytes()
}

func UnpaddingPKCS7(cipherText []byte, blockSize int) []byte {
	lastNum_byte := cipherText[len(cipherText)-1]
	lastNum := int(lastNum_byte)

	// 에러 체크 1 : 패딩이 0이면 nil 반환
	if lastNum == 0 || lastNum > len(cipherText) {
		return nil
	}

	// 에러체크 2 : 적혀있는 패딩보다 더 많이 패딩된 경우 오류로 간주하고 nil 반환
	for i := 0; i < lastNum; i++ {
		if cipherText[len(cipherText)-lastNum+i] != lastNum_byte {
			return nil
		}
	}

	// fmt.Println("Unpadding Data : ", cipherText[:len(cipherText)-lastNum])

	// 에러들을 통과하면 암호문에서 패딩들을 뺀 값을 반환
	return cipherText[:len(cipherText)-lastNum]
}

func main() {

	//32byte
	key := "01234567890123456789012345678901"
	//16byte
	iv := key[0:16]

	txt := "Hello Im fine thank you and you?"
	cTxt := AES256Encrypt(txt, key, iv)

	// XML 테스트 문자열
	// cTxt := "6asVJPcCJrjmPpl+OOZuCNNzZzLbIgtoFwyMpyVtuQNGhsgGRq5FAxRBIRpkrK5oYQg9DZ/xodwEFJFu67bFKYxXhgCMEnHbz22a9Pt2lq/+e/JJYGT2i4xm12MUxlZ7KXuxZEynbCjNaw7APwL6ZA=="

	dTxt := AES256Decode(cTxt, key, iv)
	fmt.Println("Result: ", dTxt)
}
