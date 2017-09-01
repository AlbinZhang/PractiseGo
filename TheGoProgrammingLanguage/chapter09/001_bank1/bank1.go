package main

import (
	"fmt"

	"time"

	"./bank"
)

func main() {
	b := make(chan int, 1)
	fmt.Println("-----------------")
	select {
	case b <- 122:
	default:
		fmt.Println("default-main")
	}
	time.Sleep(10 * time.Second)
	bank.Deposit(100)
	bank.Deposit(1200)
	fmt.Println(bank.Balance())
	<-b
}
