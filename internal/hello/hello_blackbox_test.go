package hello_test // note the _test suffix

import (
	"testing"

	"github.com/PhoduCoder/go-marathon/internal/hello"
)

func TestHello_BlackBox(t *testing.T) {
	got := hello.Hello("Gaurav")
	want := "Hello, Gaurav"
	if got != want {
		t.Fatalf("Hello() = %q, want %q", got, want)
	}
}
