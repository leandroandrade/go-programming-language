package main

import (
	"fmt"
	"sort"
)

func main() {
	// two forms to create a empty map
	//ages := map[string]int{}
	//ages := make(map[string]int)
	//fmt.Println(ages)

	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println(ages)
	fmt.Println(ages["alice"])

	delete(ages, "alice")
	fmt.Println(ages)

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	ages["bob"] = 20

	// sort keys of maps
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	// verify if value exists
	if age, ok := ages["bob"]; ok {
		fmt.Println("the age found", age)
	}

}
