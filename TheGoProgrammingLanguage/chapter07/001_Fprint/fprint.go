package main

import (
	"bufio"
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {

	num, _, _ := bufio.ScanWords(p, true)
	*c += ByteCounter(num)
	return len(p), nil
}

func main() {

	var count ByteCounter
	fmt.Fprintf(&count, "HelloWorld")

	fmt.Println(count)
}
