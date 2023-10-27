package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {		
	var pid string
	fmt.Print("Enter PID : ")
	_,err := fmt.Scan(&pid)
	if err!=nil {
		fmt.Println("입력 에러",err)
	}
	getPPID(pid)

	
}

func getPPID(pid string){
	cmd := exec.Command("ps", "-o", "ppid", "-p",pid)
	output,err := cmd.Output()

	defer func ()  {
		if rec := recover(); rec!=nil{
			if err.Error() == "exit status 1"{
				fmt.Println("PPID를 찾을 수 없습니다.")
			} else{
				fmt.Println("[ERROR] 명령어 에러",err)
			}
			return
		}
	}()

	if err!=nil{
		panic(err)
	}

	ppid := strings.Fields(string(output))[1]

	if ppid=="0"{
		fmt.Println("탐색을 마쳤습니다.")
		fmt.Println("프로그램을 종료합니다.")
		return
	}
	fmt.Printf("%s's PPID is %s\n",pid,ppid)

	getPPID(ppid)

}