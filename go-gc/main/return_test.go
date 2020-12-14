package main

import "testing"

var n = 1

func BenchmarkReturn1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = return1(n)
	}
}
func BenchmarkReturn2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = return2(n)
	}
}

func BenchmarkReturn3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = return3(n)
	}
}
