// Dup1 prints the text of each line that appears more than
// once in standard input, preceded by its count.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// map holds key/value pairs and allows to store, retrieve
	// or test for an item in the set
	// map[key]value
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	// don't care about the errors for now
	for line, n := range counts {
		// braces are required for the body of if statement
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
