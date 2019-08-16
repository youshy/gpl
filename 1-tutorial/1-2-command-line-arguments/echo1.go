package main

import (
	"fmt"
	// allows for getting the arguments from the command line
	"os"
)

func main() {
	// strong type declaration for s, sep
	// by default - initialized to zero value for it's type
	// for int = 0
	// for strings = ""
	var s, sep string

	// Go uses only for loops

	// the := is shorthand variable assignment - you don't have to declare the type here
	// Go will assign the type itself
	// Shorthand variable assignment is only valid inside functions

	// len - gives the length of (param)

	// lastly, i++ and i-- are statements in Go, not expressions
	// they can be used postfix only - --i wouldn't be legal in this case

	// for initialization; condition; post
	// if only condition would be used - that'd be a while loop

	// no params - infinite loop
	for i := 1; i < len(os.Args); i++ {
		// concatenates the value
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
