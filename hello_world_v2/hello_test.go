package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	if got != want {
		t.Errorf(" got %q want %q", got, want)
	}
}

func TestHelloWorld(t *testing.T) {
	assertCorrectMessage(t, HelloWorld(), "Hello, world")
}

func TestHelloWithPerson(t *testing.T) {
	t.Run("Say hello to people", func(t *testing.T) {
		assertCorrectMessage(t, HelloWithPerson("Pato"), "Hello, Pato")
	})

	t.Run("Say 'Hello, world' when an empty string is passed", func(t *testing.T) {
		assertCorrectMessage(t, HelloWithPerson(""), "Hello, world")
	})
}
