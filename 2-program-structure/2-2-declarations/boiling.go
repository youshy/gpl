// Prints the boiling point of water
package main

import "fmt"

const boilingF = 212.0 // this is package-level declaration
// will be visible throughot all the files of the package

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
