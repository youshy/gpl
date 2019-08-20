// *build arrays
package arrays

import "fmt"

func main() {
	// Arrays are less used in go, due to their fixed size

	// By default, array are initialized with zero values
	// to initialize an array with a list of values use an array literal
	// [size]type
	var q [3]int = [3]int{1, 23, 42}
	fmt.Println(q)
	// array can be also created with the length determined by the number of initializers
	r := [...]int{10, 20, 30, 40, 50}
	fmt.Printf("%T\n", r) // "[5]int"

	// there is a way to create array with k/v pairs
	type Currency int

	const (
		USD Currency = iota // this will automatically increase the indexes
		EUR
		GBP
		PLN
	)

	symbol := [...]string{USD: "\u0024", EUR: "\u20AC", GBP: "\u00A3", PLN: "PLN"}

	fmt.Println(symbol)
	fmt.Println(EUR, symbol[EUR])
	fmt.Println(USD, symbol[USD])
	fmt.Println(PLN, symbol[PLN])
	fmt.Println(GBP, symbol[GBP])

	// during initialization, if any value is unspecified, it will take zero value for the element type
	a := [...]int{99: -1} // this creates an array with 100 elements, all zero, except for the last one with "-1"
	fmt.Println(a)
}
