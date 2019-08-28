// Map is an implementation of hash table in Go
// map type is written as
// map[K]V
// where K and V are the types of its keys and values.
// All keys have to be of a same type, same for values
// But keys and values doesn't have to be the same type.

// Maps doesn't have fixed lenght

package main

import (
	"fmt"
	"sort"
)

func main() {
	Dedup()
}
func examples() {
	map1 := mapMake()
	fmt.Println(map1)
	map2 := mapLiteral()
	fmt.Println(map2)
	map3 := mapStepByStep()
	fmt.Println(map3)

	// access map element
	fmt.Println(map2["Alice"])

	// delete element
	// delete(map, element)
	delete(map3, "Tom")
	fmt.Println(map3)

	// we can access the members and modify them
	map3["Ben"]++      // increases by 1
	map2["Alice"] -= 2 // age down!
	fmt.Println(map2)
	fmt.Println(map3)

	// iterate over map
	for name, age := range map2 {
		fmt.Printf("%s\t%d\n", name, age)
	}

	map4 := map[string]int{
		"Tom":      29,
		"Jack":     30,
		"Annie":    25,
		"Harry":    13,
		"Arnold":   13,
		"Bam":      33,
		"Xeno":     99,
		"Yelp":     21,
		"Ron":      31,
		"Hermione": 11,
		"Yf":       22,
	}

	fmt.Println(map4)

	sortAndPrint(map4)

	result := equal(map[string]int{"A": 0}, map[string]int{"B": 42})
	fmt.Println(result)
}

// mapMake

func mapMake() map[string]int {
	// This creates a map with zero values using make
	// Initializes the standard types for it
	ages := make(map[string]int) // mapping from strings to ints
	return ages
}

// mapLiteral

func mapLiteral() map[string]int {
	// This creates a map with initial k/v pairs
	ages := map[string]int{
		"Alice":   31,
		"Charlie": 34,
	}
	return ages
}

// mapStepByStep

func mapStepByStep() map[string]int {
	// this combines both make and literal
	// makes a map
	ages := make(map[string]int)
	// adds new k/v pairs
	ages["Tom"] = 33
	ages["Ben"] = 19
	return ages
}

// sort map
// this is (apparently) common pattern in go
func sortAndPrint(ages map[string]int) {
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}

// comparing two maps if they contain the same k/v
// only legal comparison for maps it with nil
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
