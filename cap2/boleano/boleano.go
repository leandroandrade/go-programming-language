package main

import "fmt"

func main() {
	fmt.Println(btoi(true))
	fmt.Println(btoi(false))

	name := "Leandro!@#$%&*29"
	for _, char := range name {
		printCharNumberASCII(char)
	}
}
func printCharNumberASCII(c rune) {
	if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9' {
		fmt.Println(string(c))
	}
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
