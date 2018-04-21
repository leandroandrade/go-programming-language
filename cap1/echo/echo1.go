package main

import (
	"os"
	"fmt"
)
//agrupa os argumentos informados em Program arguments
func main() {
	var s, step string
	for i := 1; i < len(os.Args); i++ {
		s += step + os.Args[i]
		step = " "
	}
	fmt.Println(s)
}
