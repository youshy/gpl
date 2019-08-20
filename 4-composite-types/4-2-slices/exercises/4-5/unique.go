// removes duplicates in slice of strings
package unique

func unique(strings []string) []string {
	w := 0
	for _, s := range strings {
		if strings[w] == s {
			continue
		}
		w++
		strings[w] = s
	}
	return strings[:w+1]
}
