package main

import (
	"testing"
)

func BenchmarkMoveTenZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = moveZeroes(ten)
	}
}

func BenchmarkMoveHundredZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = moveZeroes(hundred)
	}
}

func BenchmarkMoveThousandZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = moveZeroes(thousand)
	}
}

func BenchmarkMove10kZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = moveZeroes(tenk)
	}
}

func BenchmarkMove100kZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = moveZeroes(hundredk)
	}
}
