package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		line := input.Text()
		counts[line]++
	}

	for k, v := range counts {
		fmt.Printf("%d\t\t%s\n", v, k)
	}
}
