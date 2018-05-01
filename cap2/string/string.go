package main

import (
	"fmt"
	"bytes"
	"strconv"
)

func main() {
	s := "hello, world"
	fmt.Println(len(s))

	fmt.Println(s[0], s[7])

	fmt.Println(s[0:5])

	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])

	fmt.Println("Hi, " + s[:])

	fmt.Println(`Exemplo usando string literais brutas.
	Uso:
		utilizar com caracteres entre crase 
	`)

	fmt.Println(string(65))
	fmt.Println(string(121212333333))
	fmt.Println(string(0x25C9))

	bytes := []byte(s)
	fmt.Println(bytes)

	fmt.Println(intsToString([]int{1, 2, 3}))

	fmt.Println(comma("12345"))
	fmt.Println(commaBuffer("12345"))

	integer:= 12
	fmt.Println(strconv.Itoa(integer))
	fmt.Println(fmt.Sprintf("%d", integer))
	fmt.Println(strconv.Atoi("12"))

}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaBuffer(s string) string {
	//buffer := &bytes.Buffer{}
	var buffer bytes.Buffer

	n := len(s) % 3

	if n == 0 {
		n = 3
	}

	buffer.WriteString(s[:n])

	for i := n; i < len(s); i += 3 {
		buffer.WriteString(",")
		buffer.WriteString(s[i : i+3])
	}
	return buffer.String()
}

//efficient to create incremental string
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
