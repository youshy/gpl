// rotates a slice of ints by one position to the left.
package main

import "fmt"

func rotateInts(ints []int) {
	first := ints[0]
	copy(ints, ints[1:])
	ints[len(ints)-1] = first
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	rotateInts(s)
	fmt.Println(s)
}
