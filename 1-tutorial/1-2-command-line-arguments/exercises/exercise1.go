package main

import (
	"fmt"
	"os"
	"strings"
)

// Print os.Args[0] as well, so print the name of the command that invoked the package

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
