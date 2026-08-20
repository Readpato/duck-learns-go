package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf(" got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("Say hello to people", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Pato", ""), "Hello, Pato")
	})

	t.Run("Say 'Hello, world' when a non recognized language is passed and world when 'name' is empty", func(t *testing.T) {
		assertCorrectMessage(t, Hello("", "French"), "Hello, world")
	})
}
