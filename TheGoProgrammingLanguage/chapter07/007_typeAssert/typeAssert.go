package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var v interface{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		v = 1
		t := r.Intn(100)
		fmt.Printf("t = %d\n", t)
		if t%2 == 0 {
			v = "hello"
		}
		if _, ok := v.(int); ok {
			fmt.Printf("%d\n", v)
		}
	}
}
