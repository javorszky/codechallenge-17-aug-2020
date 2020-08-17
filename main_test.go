package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
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

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "moves zeroes to the end correctly",
			args: args{
				nums: []int{1, 2, 0, 3, 0, 4, 5, 0, 9, 0, 3, 0, 22, 0, 41, 3, 3, 4},
			},
			want: []int{1, 2, 3, 4, 5, 9, 3, 22, 41, 3, 3, 4, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "leaves slice alone if there are no zeroes to move",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 9, 3, 22, 41, 3, 3, 4},
			},
			want: []int{1, 2, 3, 4, 5, 9, 3, 22, 41, 3, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, moveZeroes(tt.args.nums), tt.want)
		})
	}
}