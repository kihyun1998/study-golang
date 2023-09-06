package main

import (
	"fmt"
	"strconv"
)

/*
func Atoi(s string)(i int, err error) : 숫자로 이루어진 문자열을 숫자로 변환 [ string > int ]

func Itoa(i int) string : 숫자를 문자열로 변환 [ int > string ]

func FormatBool(b bool) string : 불 값을 문자열로 변환 [ bool > string ]

func FormatFloat(f float64, fmt byte, prec, bitSize int) string : 실수를 문자열로 변환 [ float64 > string ]
2번째 parameter의 argument 종류
- 'b' : 2진수 지수
- 'e','E' : 10진수 지수
- 'x', 'X' : 16진수 가수와 2진수 지수
- 'f' : 지수 없음
- 'g' : 큰 지수의 경우

ex) s := strconv.FormatFloat(3.141592, 'E', -1, 64)

func FormatInt(i int64, base int) string : 부호 있는 정수를 문자열로 변환

func FormatUint(i uint64, base int) string : 부호 없는 정수를 문자열로 변환

func AppendBool(dst []byte, b bool)[] byte : 불 값을 문자열로 변환하여 슬라이스 뒤에 추가

func AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int)[]byte : 실수를 문자열로 변환하여 슬라이스 뒤에 추가

func AppendInt(dst []byte, i int64, base int)[]byte : 부호 있는 정수를 문자열로 변환하여 슬라이스 뒤에 추가

.
.
.
*/

func main() {
	num1, err := strconv.Atoi("100")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num1)

	num2, err := strconv.Atoi("10c")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num2)

	s := strconv.Itoa(num1)
	fmt.Println(s)
}
