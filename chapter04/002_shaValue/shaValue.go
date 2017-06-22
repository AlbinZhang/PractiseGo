package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var shaType = flag.String("s", "256", "256 - sha256\n384 - sha384\n512 - sha512")

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("eg: go run shaValue.go test")
		os.Exit(1)
	}
	buf := os.Args[1]
	if os.Args[1] == "-s" {
		if len(os.Args) <= 3 {
			os.Exit(2)
		}
		buf = os.Args[3]
	}

	switch *shaType {
	case "256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(buf)))
	case "384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(buf)))
	case "512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(buf)))
	default:
		fmt.Println("sha type just support 256 384 512")
	}
}
