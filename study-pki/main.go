package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"time"
)

type VerificationPackage struct{

}


// 디지털 인증서
type DigitalCertification struct{
	IssuerInfo string
	OwnerInfo string
	Authority string
	//발급자의 전자서명, 유효기간은 일단 뺐다.
	OwnerPublicKey *rsa.PublicKey
}

// 인증서 발급
func (di *DigitalCertification)CertificationRegister(name,authority string, ownerPubKey *rsa.PublicKey){
	di.IssuerInfo = "CA"
	di.OwnerInfo = name
	di.Authority = authority
	di.OwnerPublicKey = ownerPubKey
}



// 키 저장
type CA_Keys struct{
	privateKey *rsa.PrivateKey
	publicKey *rsa.PublicKey
}


// 키 로컬 저장 예제 : https://gist.github.com/sdorra/1c95de8cb80da31610d2ad767cd6f251
func (keys *CA_Keys)CA_make_key() {
	privateKey,err:=rsa.GenerateKey(rand.Reader, 2048)
	if err != nil{
		fmt.Println("[ERROR] 암호화",err)
		return
	}
	publicKey := &privateKey.PublicKey

	keys.privateKey = privateKey
	keys.publicKey = publicKey
}

// SHA256 해시함수
func HashMessage(message string) []byte {
	hash := sha256.New()
	hash.Write([]byte(message))
	md := hash.Sum(nil)
	// 문자열 반환 시
	// mdStr := hex.EncodeToString(md)

	return md
}

func main() {

	var want string

	fmt.Println("===============================")
	fmt.Println("|         **Welcome**         |")
	fmt.Println("===============================")
	fmt.Println("| This is PKI Example Program |")
	fmt.Println("===============================")

	fmt.Println("\n\n ")
	fmt.Println("안녕하세요 인증서를 발급하시겠습니까? [y/N]")
	
	_,err := fmt.Scan(&want)
	if err!=nil{
		fmt.Println("[ERROR] 입력",err)
	}

	for {
		if want == "y" || want == "N" {
			break
		}else{
			fmt.Println("y 또는 N만 입력해주세요.")
		}
		_,err := fmt.Scan(&want)
		if err!=nil{
			fmt.Println(err)
		}
	}

	// 종료
	if want=="N"{
		return
	}

	// 0. 사용자 정보 얻기
	name := "park"
	authority := "Root"


	// 1. CA로 키 발급
	keys := CA_Keys{}
	keys.CA_make_key()

	fmt.Println("\n\nkey 가 발급되었습니다.")

	// 2. 인증서 발급
	di := DigitalCertification{}
	di.CertificationRegister(name,authority,keys.publicKey)

	for i:=0;i<3;i++{
		fmt.Println(".")
		time.Sleep(700 * time.Millisecond)	
	}
	
	fmt.Println("인증서가 발급됐습니다.")

	// 3. 디지털 서명 생성
	message := "I am Park Ki Hyun"
	hashMessage := HashMessage(message)
	
	signature,err := rsa.SignPKCS1v15(
		rand.Reader,
		keys.privateKey,
		crypto.SHA256, 
		hashMessage,
	)
	if err!=nil {
		fmt.Println("[ERROR] 전자서명: ",err)
	} 
}

