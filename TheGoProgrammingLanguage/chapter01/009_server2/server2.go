// go run server2.go
// http://127.0.0.1:8000/echo
// http://1270.0.01:8000/count

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

var address = flag.String("address", ":10001", "http address")

func main() {
	flag.Parse()

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/heart", heart)
	log.Fatal(http.ListenAndServe(*address, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count = %d", count)
	mu.Unlock()
}

func heart(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
