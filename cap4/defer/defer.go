package main

import (
	"time"
	"log"
)

func main() {
	slowOperation()
}

func slowOperation() {
	defer trace("slowOperation")() // parenteses necessario, se nao a funcao executa somente o inicio.

	time.Sleep(2 * time.Second)
}

func trace(message string) func() {
	start := time.Now()
	log.Printf("enter %s", message)

	return func() {
		log.Printf("exit %s (%s)", message, time.Since(start))
	}
}
