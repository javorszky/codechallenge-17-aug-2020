package main

import (
	"fmt"
	"math/rand"
)

func main() {
	input := []int{1, 2, 0, 1, 0, 0, 3, 6}
	output := moveZeroes(input)

	fmt.Printf("%v\n", output)

	lenghts := []uint32{10, 100, 1000, 10000, 100000}
	for _, v := range lenghts {
		fmt.Printf("%#v\n\n", generateSlice(v))
	}
}

func moveZeroes(nums []int) []int {
	moved := make([]int, 0, len(nums))
	zeroes := 0
	for _, v := range nums {
		if v == 0 {
			zeroes++
			continue
		}
		moved = append(moved, v)
	}
	for i := 0; i < zeroes; i++ {
		moved = append(moved, 0)
	}

	return moved
}

func generateSlice(length uint32) []int {
	nums := make([]int, 0, length)
	for i := uint32(0); i < length; i++ {
		nums = append(nums, rand.Intn(36))
	}
	return nums
}
