package iteration

import "strings"

func Repeat(input string, n int) string {
	var repeated strings.Builder

	for i := 0; i < n; i++ {
		repeated.WriteString(input)
	}

	return repeated.String()
}
