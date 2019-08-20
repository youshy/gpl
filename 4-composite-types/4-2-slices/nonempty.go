// Example of an in-place slice algorithm
package main

// returns a slice holding only the non-empty strings
// method is destructive, the array is modified during the call
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
