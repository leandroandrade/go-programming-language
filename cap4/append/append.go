package main

import "fmt"

func main() {
	var numbers []int
	numbers = append(numbers, 1)
	fmt.Println(numbers)

	numbers = append(numbers, 2, 3)
	fmt.Println(numbers)

	fmt.Println("top is:", numbers[len(numbers)-1])

	//pop
	numbers = numbers[:len(numbers)-1]
	fmt.Println("remove element", numbers)

	var values []int
	values = append(values, 5, 6, 7, 8, 9)
	fmt.Println(values)
	fmt.Println(values[2:])
	fmt.Println(values[3:])

	// values[2:] = 7, 8, 9
	// values[3:] = 8, 9
	// result     = 8, 9, 9
	copy(values[2:], values[3:])
	fmt.Println(values)
	fmt.Println(values[:len(values) - 1])
}
