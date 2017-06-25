package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	log.Println("_start")
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	log.SetPrefix("wait: ")
	log.SetFlags(0)
	if err := WaitForServer(os.Args[1]); err != nil {
		log.Fatal(err)
	}
	fmt.Println("successful")
}
