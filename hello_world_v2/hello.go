package main

import (
	"fmt"
	"strings"
)

const (
	spanishLanguage = "spanish"
	germanLanguage  = "german"

	germanHelloPrefix  = "Hallo, "
	spanishHelloPrefix = "Hola, "
	englishHelloPrefix = "Hello, "
)

func getGreetingPrefix(language string) (prefix string) { // Cool other way to create and return a variable!
	switch language {
	case spanishLanguage:
		prefix = spanishHelloPrefix
	case germanLanguage:
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string { // Love the params type assignment like this
	lowerCaseLanguage := strings.ToLower(language) // Cool that you can declare variables like this You don't need to asign a let or const
	if name == "" {
		switch lowerCaseLanguage {
		case "spanish":
			name = "mundo"
		case "german":
			name = "welt"
		default:
			name = "world"
		}
	}

	return getGreetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("Pato", "Spanish"))
}
