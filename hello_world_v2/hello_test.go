package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := Hello()
	want := "Hello, world!"

	if got != want {
		t.Errorf(" got %q want %q", got, want)
	}
}

func TestHelloWithGreeting(t *testing.T) {
	got := HelloWithGreeting("Pato")
	want := "Hello, Pato!"

	if got != want {
		t.Errorf(" got %q want %q", got, want)
	}
}
