package main

import (
	"fmt"
)

func main() {
	// slices are variable-length sequences with elements of the same type thorough
	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	endlessSummer := summer[:5] // this extends the "summer" slice
	fmt.Println(endlessSummer)

	// reversing the array
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	fmt.Println(a)
	reverse(a[:])
	fmt.Println(a)

	var runes []rune // creates a slice of runes
	str := "I will find my 周边"
	for _, r := range str {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	// working with appendInt
	fmt.Printf("\nworking with appendInt\n")
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d	cap=%d\n%v\n", i, cap(y), y)
		x = y
	}
	// or using built-in append
	fmt.Printf("\nUsing built-in append\n")
	var newX []int
	newX = append(newX, 1)
	newX = append(newX, 2,3,4,5)
	newX = append(newX, newX...) // append the slice x
	fmt.Println(newX)

	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
}
