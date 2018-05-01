package main

import "fmt"

//numbers withou signal(uint) cannot receive negative numbers
func main() {
	var i int8 = -2

	//try to put a negative signal and execute the program
	var j uint8 = 2

	fmt.Println(i, j)
}
