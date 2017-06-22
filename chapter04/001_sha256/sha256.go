package main

import "fmt"
import "crypto/sha256"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥", 12345: "123213213"}
	fmt.Println(RMB, symbol[RMB])
	fmt.Printf("%T\n", symbol)

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Printf("diff count = %d\n", diffByteCount(c1, c2))
}

func diffByteCount(s1 [sha256.Size]byte, s2 [sha256.Size]byte) int {
	var count int
	for i, v := range s1 {
		fmt.Println(v, s2[i])
		if v != s2[i] {
			count++
		}
	}
	return count
}
