package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 200
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToc(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g F", f)
}

func main() {
	c := FToc(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)

	fmt.Printf("%s\n", c)
	fmt.Println(c)

	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))



}
