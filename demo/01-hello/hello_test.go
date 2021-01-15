package hello_test

import (
	"hello"
	"testing"
)

func TestHello(t *testing.T) {
	result := hello.SayHi("somkiat")
	if result != "Hello somkiat" {
		t.Errorf("Error %v", result)
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hello.SayHi("somkiat")
	}
}
