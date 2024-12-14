package main

import (
	"cmp"
	"fmt"
	"math"
)

func main() {
	fmt.Println("hey")
	ints := []int{0, 1, 2, 3, 4}
	strs := []string{"0", "1", "2", "3", "4"}
	bols := []bool{false, false, false, true, true}

	assert(linearSearch(ints, 4), true, "linear search found failed")
	assert(linearSearch(strs, "4"), true, "linear search found failed")
	assert(linearSearch(ints, 10), false, "linear search not found failed")
	assert(linearSearch(strs, "10"), false, "linear search not found failed")

	assert(binarySearch(ints, 0), true, "binary search found failed")
	assert(binarySearch(ints, 1), true, "binary search found failed")
	assert(binarySearch(ints, 2), true, "binary search found failed")
	assert(binarySearch(ints, 3), true, "binary search found failed")
	assert(binarySearch(ints, 4), true, "binary search found failed")
	assert(binarySearch(ints, 10), false, "binary search not found failed")
	assert(binarySearch(ints, -1), false, "binary search not found failed")

	assert(crystalBalls(bols), 3, "crystal balls failed")
}

// Given two crystal balls that will break if dropped from high enough
// distance, determine the exact spot in which it will break in the most
// optimized way.
func crystalBalls(arr []bool) int {
	jump := int(math.Sqrt(float64(len(arr))))

	i := jump
	for ; i < len(arr); i += jump {
		if arr[i] == true {
			break
		}
	}

	i -= jump
	for ; i < i+jump && i < len(arr); i++ {
		if arr[i] == true {
			return i
		}
	}

	return -1
}

// Array needs to be sorted.
func binarySearch(arr []int, v int) bool {
	return binarySearchRec(arr, 0, len(arr)-1, v)
}

func binarySearchRec[T cmp.Ordered](arr []T, lowIdx int, highIdx int, v T) bool {
	// fmt.Println(arr[lowIdx:highIdx+1], lowIdx, highIdx, v)

	// single element left
	if lowIdx == highIdx {
		return arr[lowIdx] == v
	}

	// two elements left
	if lowIdx+1 == highIdx {
		return arr[lowIdx] == v || arr[highIdx] == v
	}

	midIdx := (highIdx - lowIdx + 1) / 2

	// found element in the middle
	if arr[midIdx] == v {
		return true
	}

	// search lower half
	if v < arr[midIdx] {
		return binarySearchRec(arr, lowIdx, midIdx-1, v)
	}

	// search high half
	return binarySearchRec(arr, midIdx+1, highIdx, v)
}

func linearSearch[T comparable](arr []T, v T) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == v {
			return true
		}
	}
	return false
}

func assert[T comparable](a T, b T, msg string) {
	if a != b {
		compare := fmt.Sprintf("a = %v\nb = %v", a, b)
		panic(fmt.Sprintf("Assert Error: %s\n\n%s", msg, compare))
	}
}
