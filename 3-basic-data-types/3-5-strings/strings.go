// *build strings

// strings in go are immutable
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("BASICS\n")
	s := "hello, world"
	fmt.Println(len(s))     // gives the lenght
	fmt.Println(s[0], s[7]) // prints out i-th byte of string (not the character)

	// substring
	// s[i:j]
	fmt.Println("\nSUBSTRINGS\n")
	fmt.Println(s[0:5]) // "hello"
	// both i or j can be ommited, then
	// i = 0
	// j = len(string)
	fmt.Println(s[:5]) // "hello"
	fmt.Println(s[7:]) // "world"
	fmt.Println(s[:])  // "hello, world"

	// string literals
	fmt.Println("\nSTRING LITERALS\n")
	fmt.Println("\xe4\xb8\x96\xe7\x95\x8c") // per byte encoding
	fmt.Println("\u4e16\u754c")             // 16-bit unicode value
	fmt.Println("\U00004e16\U0000754c")     // 32-bit unicode value

	newS := "Hello, \u4e16\u754c"
	fmt.Println("\nDecode UTF-8 string")
	for i, r := range newS {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	fmt.Println("count runes")
	fmt.Println(utf8.RuneCountInString(newS))

	fmt.Println("katakana")
	kata := "\u30d7\u30ed\u30b0\u30e9\u30e0"
	fmt.Println(kata)
	fmt.Println("to hex")
	fmt.Printf("%x\n", kata)
	r := []rune(kata)
	fmt.Println("to unicode")
	fmt.Printf("%x\n", r)

	// to modify any string, you have to cast it into []byte (byte slice)
	// Strings can be freely converted to byte slices and back again

	str := "abc"
	byt := []byte(str)
	str2 := string(byt)
}
