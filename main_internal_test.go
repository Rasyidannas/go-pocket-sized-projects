package main

import "testing"

func TestGreet(t *testing.T) {
	want := "Hello world"

	got := greet()

	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}
