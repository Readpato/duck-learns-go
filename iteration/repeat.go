package iteration

import "strings"

// Super cool that you can import a specific package for strings
// and you get a lot of functionality to work with the machine itself
// The Omnissiah would be proud

func Repeat(character string, count int) string {
	var repeated strings.Builder
	for i := 0; i < count; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
