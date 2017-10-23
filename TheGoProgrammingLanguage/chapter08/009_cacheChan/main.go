package main

import (
	"fmt"
	"log"
)

func mirroredQuery() string {
	responses := make(chan string, 1)
	go func() { log.Println("1"); responses <- "33333333333333333" }()
	go func() { log.Println("2"); responses <- "22222222222222222" }()
	go func() { log.Println("3"); responses <- "11111111111111111" }()
	return <-responses
}

func main() {
	fmt.Println(mirroredQuery())
	fmt.Println(mirroredQuery())
	fmt.Println(mirroredQuery())
}
