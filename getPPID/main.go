package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

var once sync.Once

type singleton map[int]int

var globalMap singleton

// 싱글톤 객체 생성
func NewMap() singleton {
	// 한번만 실행
	once.Do(func() {
		globalMap = make(singleton)
	})

	return globalMap
}

func AddValueByKey(key int, value int) {
	globalMap := NewMap()
	globalMap[key] = value
}

func ShowItem() {
	globalMap := NewMap()
	for k, v := range globalMap {
		fmt.Println("PID : ", k, "\tPPID : ", v)
	}
}

func GetLength() int {
	globalMap := NewMap()
	return len(globalMap)
}

func main() {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("프로그램을 종료합니다.")
			return
		}
	}()

	var sPID string

	fmt.Print("Enter PID : ")
	_, err := fmt.Scan(&sPID)
	if err != nil {
		fmt.Println("숫자만 입력해주세요.")
		panic(err)
	}

	PID, err := strconv.Atoi(sPID)
	if err != nil {
		fmt.Println("[ERROR]숫자를 입력해주세요.")
		panic(err)
	}

	getPPID(PID)
	if mapLen := GetLength(); mapLen == 0 {
		fmt.Println("PPID를 찾을 수 없습니다.")
		fmt.Println("프로그램을 종료합니다.")
		return
	} else {
		ShowItem()
		fmt.Println("탐색을 마치고 프로그램을 종료합니다.")
	}
}

func getPPID(PID int) {

	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("PPID를 찾을 수 없습니다.")
			return
		}
	}()

	procStatus := fmt.Sprintf("/proc/%d/status", PID)
	cmd := exec.Command("cat", procStatus)

	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	ppidInOutput := strings.Split(string(output[:]), "\n")
	ppidStr := ppidInOutput[6]
	sPPID := strings.Fields(ppidStr)[1]
	PPID, _ := strconv.Atoi(sPPID)
	if PPID == 0 {
		return
	} else {
		AddValueByKey(PID, PPID)
		getPPID(PPID)
	}
}
