package main

import "fmt"

func main() {
	multipleSum()
}

func multipleSum() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(3, 4, 2, 5, 6, 2))
	fmt.Println(sum(10, 20, 30, 40))

	// vorks also with an array of int
	values := []int{1, 2, 3, 4}
	// the <variable>... syntax spreads the values for the function
	fmt.Println(sum(values...))
}

// this function takes multiple int values
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

// basically, variadic functions are functions
// that allows for passing unlimited values
// to the function
