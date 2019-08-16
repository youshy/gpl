package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// that's the fastest way and the smallest garbage-wise
	fmt.Println(strings.Join(os.Args[1:], " "))
}
