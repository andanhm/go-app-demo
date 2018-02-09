package main

import "testing"

func BenchmarkSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sum(10, n)
	}
}
