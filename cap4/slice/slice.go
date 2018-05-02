package main

import "fmt"

func main() {
	var names []string
	fmt.Println(names)

	names = append(names, "duke")
	fmt.Println(names)

	names = append(names, "john")
	fmt.Println(names)

	ages := make([]int, 0)
	fmt.Println(ages)

	ages = append(ages, 20)
	fmt.Println(ages)

	ages = append(ages, 10)
	fmt.Println(ages)


}
