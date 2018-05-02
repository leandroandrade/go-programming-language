package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	symbol := [...] string{USD: "$", EUR: "E", GBP: "F", RMB: "Y"}
	//symbol := [] string{USD: "$", EUR: "E", GBP: "F", RMB: "Y"}
	fmt.Println(symbol[USD])

	numbers := []int{1, 2, 3}
	fmt.Println(numbers)

	names := [3]string{"paul", "john", "may"}
	fmt.Printf("%T\n", names)
}
