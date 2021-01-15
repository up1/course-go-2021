package demo_test

import (
	"demo"
	"testing"
)

func BenchmarkSolution01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		demo.StructToJson01()
	}
}

func BenchmarkSolution02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		demo.StructToJson02()
	}
}
