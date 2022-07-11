package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Umar")
	want := "Hello, Umar"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
