// This function is specialized in appending new Int to []int (integer slices)
package main

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// Still can increase, therefore - extend the slice
		z = x[:zlen]
	} else {
		// There is insufficient space - allocate a new array
		// Grow by doubling for amortized linear complexity
		// in other words - when the array will be too small
		// it creates a new array with double the size
		// and copies ints from previous one
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap) // make is a built-in function for initializing a slice. Requires []T, len, cap. cap can be ommited - then cap = len
		copy(z, x)                  // built-in function
	}
	z[len(x)] = y
	return z
}
