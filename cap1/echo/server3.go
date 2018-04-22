package main

import (
	"net/http"
	"log"
	"fmt"
	"strconv"
)

func main() {
	http.HandleFunc("/", handlerHeader)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handlerHeader(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "%s %s %s\n", request.Method, request.URL, request.Proto)
	for k, v := range request.Header {
		fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(writer, "Host = %q\n", request.Host)
	fmt.Fprintf(writer, "RemoteAddr = %q\n", request.RemoteAddr)

	if err := request.ParseForm(); err != nil {
		log.Print(err)
	}

	fmt.Println(strconv.Atoi(request.URL.Query().Get("page")))
	fmt.Println(strconv.Atoi(request.URL.Query().Get("size")))

	for k, v := range request.Form {
		fmt.Fprintf(writer, "Form[%q] = %q\n", k, v)
	}
}
