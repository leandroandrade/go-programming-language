package main

import (
	"bufio"
	"strings"
	"fmt"
)

func main() {
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	frequency := make(map[string]int)

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	// Count the words.
	for scanner.Scan() {
		frequency[string(scanner.Bytes())]++
	}

	for word, freq := range frequency {
		fmt.Printf("%s\t%d\n", word, freq)
	}

}
