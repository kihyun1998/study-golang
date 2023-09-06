package legacy

// 기본 ==========================
// func hello(n int) {
// 	r := rand.Intn(100)
// 	time.Sleep(time.Duration(r))
// 	fmt.Println(n)
// }

// func main() {
// 	for i := 0; i < 100; i++ {
// 		go hello(i)
// 	}
// 	fmt.Scanln()
// }

// 멀티코어 ========================
// func main() {
// 	runtime.GOMAXPROCS(runtime.NumCPU()) //멀티코어 사용
// 	fmt.Println(runtime.GOMAXPROCS(0))

// 	s := "hello"

// 	for i := 0; i < 100; i++ {
// 		go func(n int) {
// 			fmt.Println(s, n)
// 		}(i) // << 얘가 함수 인자값으로 들어감
// 	}
// 	fmt.Scanln()
// }

/*
go routine에서

go func(입력){
	명령어
}(인자값)

이다.
*/
