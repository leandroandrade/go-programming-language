package main

import (
	"net/http"
	"log"
	"sync"
	"fmt"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func counter(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	fmt.Fprintf(writer, "Count %d\n", count)
	mu.Unlock()
}
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("acessou handler")
	mu.Lock()
	count++
	mu.Unlock()

	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
}
