package main

import (
	"flag"
	"crypto/sha256"
	"crypto/sha512"
	"os"
	"fmt"
	"bufio"
	"errors"
	"log"
)

var crypt = flag.Int("c", 256, "type of crypto length(256 or 512)")

func parser(value []byte) ([]byte, error) {
	switch *crypt {
	case 256:
		encrypted := sha256.Sum256(value)
		return encrypted[:], nil
	case 512:
		encrypted := sha512.Sum512(value)
		return encrypted[:], nil

	default:
		return nil, errors.New("invalid crypto type")
	}

}

func main() {
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter with value to parse to %v: ", *crypt)
	value, _ := reader.ReadString('\n')

	sha, err := parser([]byte(value))
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%x\n", sha)
}
