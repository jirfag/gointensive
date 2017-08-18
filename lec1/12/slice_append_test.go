package main

import "testing"

func BenchmarkSliceAppending(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genRandomNums(10000)
	}
}

func BenchmarkSliceAppendingOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genRandomNumsOptimized(10000)
	}
}
