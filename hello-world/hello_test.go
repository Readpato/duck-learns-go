package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

	})

	t.Run("Say 'Hello, world' when there's no provided name", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, world' in Spanish", func(t *testing.T) {
		got := Hello("Natalia", "spanish")
		want := "Hola, Natalia"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, world' in French", func(t *testing.T) {
		got := Hello("Pedro", "french")
		want := "Bonjour, Pedro"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
