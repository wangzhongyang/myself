package main

import "testing"

func Benchmark_NewByPool(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := NewByPool()
		_ = p.String()
		p.Name = "name1"
		PeoplePool.Put(p)
	}
}

func Benchmark_NeyPeople(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := NeyPeople()
		_ = p.String()
	}
}
