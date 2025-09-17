package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Gaurav")
	want := "Hello, Gaurav"

	if got != want {
		t.Fatalf("Hello() = %q; want %q", got, want)
	}
}
