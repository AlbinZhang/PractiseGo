package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}

	return
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("loading")
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
}
