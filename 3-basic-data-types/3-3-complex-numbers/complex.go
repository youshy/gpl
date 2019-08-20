// *build complex
package main

import "fmt"

func main() {
	// there are two sizes of complex numbers in go
	// complex64 (made of float32)
	// complex128 (made of float64)
	// contains build-in function 'complex' creates a complex number from both real and imaginary components
	// and the build-in `real` and `imag` extracts those components

	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x * y)
	fmt.Println(real(x))
	fmt.Println(real(y))
	fmt.Println(imag(x))
	fmt.Println(imag(y))
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))

	// go allows for simplyfying x and y syntax
	newX := 1 + 2i
	newY := 3 + 4i
}
