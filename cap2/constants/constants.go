package main

import "fmt"

type Weekday int

const(
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednsday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
}
