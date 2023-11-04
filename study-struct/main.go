package main

import (
	"fmt"
	"study-golang/study-struct/accounts"
)

func main() {
	account := accounts.NewAccount("park")
	account.Deposit(100)

	fmt.Println("account : ", account)
}
