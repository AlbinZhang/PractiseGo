package main

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

const debug = false

func main() {

	fmt.Println(time.Now().Unix())
	//var buf *bytes.Buffer
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
	if debug {
		fmt.Println("main debug = true")
	}
}

func f(out io.Writer) {
	// ...do something...
	fmt.Printf("%T, %v\n", out, out)
	if out != nil {
		fmt.Println("000000000")
		out.Write([]byte("done!\n"))
	}
}
