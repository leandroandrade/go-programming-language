package main

import "fmt"

func main() {
	fmt.Println(returnValue())
}

func returnValue() (value string) {
	defer func() {
		recover()
		value = "few"
	}()

	panic("omg")
}
