package main

import (
	"fmt"
	"os"
	"strconv"

	tempconv "github.com/gpl/2-program-structure/2-6-packages-and-files"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)

		fmt.Printf("%s = %s, %s = %s\n%s = %s, %s = %s\n%s = %s, %s = %s\n", c, tempconv.CToF(c), c, tempconv.CToK(c), f, tempconv.FToC(f), f, tempconv.FToK(f), k, tempconv.KToF(k), k, tempconv.KToC(k))
	}
}
