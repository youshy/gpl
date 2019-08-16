package main

import (
	"fmt"
	"os"
)

// Print the inved and value of each of argument, one per line

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Printf("Index %s; arg %s\n", index, arg)
	}
}
