// Converts Celsius and Fahrenheit both ways
package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// It's possible to define new behaviors for the value of the type
// it's called type's methods in Go

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {
	fmt.Printf("%g °C\n", BoilingC-FreezingC) // 100 °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g °F\n", boilingF-CToF(FreezingC)) // 180 °F
	// fmt.Printf("%g °C", boilingF-FreezingC) // compile error: type mismatch

	// you can't compare two different types
	// fmt.Println(c == f) // compile error: type mismatch

	// using type's methods
	c := FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"; does not call String
	fmt.Println(float64(c)) // "100"; does not call String
}
