package main

import (
	"flag"
	"fmt"
	"strings"
)

// Using pointers with flag package which uses a command line arguments
// to set the values of certain variables
// distributed through program

var n = flag.Bool("n", false, "omit trailing newline") // causes echo to omit the trailing newline
// this function takes the name of a flag, variable default value and a message when using -h or -help
var sep = flag.String("s", " ", "separator") // causes echo to separate the output arguments

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
