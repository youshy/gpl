package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// because range produces a pair of values - the index and the value of the element at that index
	// we need to deal with both values
	// _ acts as a blank identifier. it assigns the value, but never reads from it
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
