package main

import (
	"fmt"
	"runtime"
	"os"
)

func main() {
	defer printStack()
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x-1)
	defer fmt.Printf("defer %d\n", x)

	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)

	fmt.Println("print into stdout")
	fmt.Println("\n\nerror\n\n", string(buf[:n]))
	os.Stdout.Write(buf[:n])
}
