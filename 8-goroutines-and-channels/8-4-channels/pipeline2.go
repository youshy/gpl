package main

import "fmt"

// aka pipeline with fixed length and not breaking
// when the channel ends

func main() {
	fixedPipeline()
}

func fixedPipeline() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer
	for x := range squares {
		fmt.Println(x)
	}
}
