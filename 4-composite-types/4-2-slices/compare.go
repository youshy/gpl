// Unlike arrays, slices cannot be compared to each other, even if they have the same elements
// if comparing two slices of bytes, you can use bytes.Equal
// for any other, you need to implement the function yourself
package main

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
