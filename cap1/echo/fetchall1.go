package main

import (
	"time"
	"os"
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
)

// busca URLs em paralelo e informa os tempos gastos e o tamanho
func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetchall(url, ch)
	}

	for range os.Args[1:] {
		write(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchall(url string, ch chan<- string) {
	start := time.Now()

	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func write(time string) {
	fmt.Println("passou aqui")
	arquivo, err := os.OpenFile("time-log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Errorf("erro %v", err)
	}

	defer arquivo.Close()

	arquivo.WriteString(time + "\n")
}
