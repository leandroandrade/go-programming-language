package main

import (
	"os"
	"net/http"
	"fmt"
	"io"
	"strings"
)

// exibe o conteudo encontrado em cada URL informada usando io.Copy, o que melhora a performance
func main() {
	protocol := "http://"

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, protocol) {
			url = protocol + url
		}

		response, err := http.Get(url)
		fmt.Printf("http status code: %v\n", response.StatusCode)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// improve performance using io.Copy
		byte, err := io.Copy(os.Stdout, response.Body)
		response.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}
		fmt.Printf("%s", byte)
	}
}
