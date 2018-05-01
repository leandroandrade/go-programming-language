package main

import (
	"flag"
	"crypto/sha256"
	"crypto/sha512"
	"log"
	"os"
	"fmt"
	"bufio"
)

var crypto = flag.Int("c", 256, "type of crypto length(256 or 512)")

func main() {
	flag.Parse()
	var parser func(value []byte) []byte

	switch *crypto {
	case 256:
		parser = func(value []byte) []byte {
			encrypted := sha256.Sum256(value)
			return encrypted[:]
		}
	case 512:
		parser = func(value []byte) []byte {
			encrypted := sha512.Sum512(value)
			return encrypted[:]
		}
	default:
		log.Fatal("Invalid crypto type")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter with value to parse to %v: ", *crypto)
	value, _ := reader.ReadString('\n')

	fmt.Printf("%x\n", parser([]byte(value)))
}
