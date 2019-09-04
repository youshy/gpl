package main

import (
	"bufio"
	"fmt"
	"strings"
)

type ByteCounter int

func main() {
	var c ByteCounter
	c.Write([]byte("Hello"))
	fmt.Println(c)
	c = 0
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "Hello, %s!", name)
	fmt.Println(c)
	var words = "My words are words of words from worlds"
	c = 0
	c.CountWords([]byte(words))
	fmt.Printf("%s:\t%v\n", words, c)
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

// ex 7.1
// count words
func (c *ByteCounter) CountWords(p []byte) (int, error) {
	s := string(p)
	scanner := bufio.NewScanner(strings.NewReader(s)) // accepts byte stream as input
	scanner.Split(bufio.ScanWords)
	// counting words
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += ByteCounter(count)
	return count, nil
}
