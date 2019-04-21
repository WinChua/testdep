package testdep

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, World"
    if got := DepHello(); got != want {
        t.Fatal("Hello", got)
    }
}
