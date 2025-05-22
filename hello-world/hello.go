package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "spanish"
const french = "french"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := englishHelloPrefix

	switch language {
		case spanish:
			prefix = spanishHelloPrefix
		case french:
			prefix = frenchHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
