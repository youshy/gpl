package main

import "fmt"

func main() {
	f := squares()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// A function literal is written like a functioon declaration
// but without a name following the func keyword

// squares returns a function that returns
// the next square number each time it's called
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
