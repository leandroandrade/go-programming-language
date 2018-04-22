package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

	r := newInt()
	s := newIntDummy()
	fmt.Println(r == s)

}

func newInt() *int {
	return new(int)
}

func newIntDummy() *int {
	var dummy int
	return &dummy
}
