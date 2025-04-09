package iteration

import "strings"

func Repeat(input string) string {
	var repeated strings.Builder

	for i := 0; i < 5; i++ {
		repeated.WriteString(input)
	}

	return repeated.String()
}
