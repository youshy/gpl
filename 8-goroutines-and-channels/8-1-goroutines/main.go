package main

import (
	"fmt"
	"time"
)

func main() {
	// this goroutine runs the spinner function
	// and then allows for further execution
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow for a reason
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

// spinner alterates through range
// to create simple spinner
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			// \r is carriage return
			// returns to the line beginning
			// without advancing to the new one
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

// fib is an implementation of fibonacci sequence
// with a recursive method
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
