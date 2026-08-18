package main

import (
	"fmt"
	"strings"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const germanHelloPrefix = "Hallo, "

func HelloWorld() string {
	return englishHelloPrefix + "world"
}

func HelloWithPerson(name string) string {
	if name == "" {
		name = "world"
	}

	return englishHelloPrefix + name

}

func HelloWithPersonAndLanguage(name, language string) string {
	lowerCaseLanguage := strings.ToLower(language)
	if name == "" {
		switch lowerCaseLanguage {
		case "spanish":
			name = "mundo"
		case "german":
			name = "welt"
		}
	}
	switch lowerCaseLanguage {
	case "spanish":
		return spanishHelloPrefix + name
	case "german":
		return germanHelloPrefix + name
	}
	return englishHelloPrefix + name

}

func main() {
	fmt.Println(HelloWorld())
}
