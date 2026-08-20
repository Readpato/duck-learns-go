package main

import (
	"fmt"
	"strings"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const germanHelloPrefix = "Hallo, "
const spanishLanguage = "spanish"
const germanLanguage = "german"

func Hello(name, language string) string {
	lowerCaseLanguage := strings.ToLower(language)
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
	switch lowerCaseLanguage {
	case spanishLanguage:
		return spanishHelloPrefix + name
	case germanLanguage:
		return germanHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Pato", "Spanish"))
}
