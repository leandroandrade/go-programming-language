package main

import (
	"fmt"
	"github.com/leandroandrade/go-programming-language/cap2/tempconv/base"
)

func init() {
	fmt.Println("initializing main")
}

func main() {
	fmt.Println(tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
}
