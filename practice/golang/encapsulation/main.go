package main

import (
	"encapsulation/users"
	"fmt"
)

func main() {
	acc := users.NewAccount("John")
	acc.Deposit(100)
	fmt.Println(acc.BalanceCount())
}
