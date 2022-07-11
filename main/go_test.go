package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	autoCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to the people", func(t *testing.T) {
		got := Hello("Umar", "")
		want := "Hello, Umar"
		autoCorrectMessage(t, got, want)

	})

	t.Run("say Hello, World when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		autoCorrectMessage(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		autoCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Pia", "French")
		want := "Bonjour, Pia"
		autoCorrectMessage(t, got, want)
	})

}
