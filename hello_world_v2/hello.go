package main

import "fmt"

const englishHelloPrefix = "Hello, "

func HelloWorld() string {
	return englishHelloPrefix + "world"
}

func HelloWithPerson(name string) string {
	if name == "" {
		name = "world"
	}

	return englishHelloPrefix + name

}

func main() {
	fmt.Println(HelloWorld())
}
