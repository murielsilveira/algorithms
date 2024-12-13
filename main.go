package main

import (
	"fmt"
)

func main() {
	fmt.Println("hey")
	arr := []int{1, 2, 3, 4, 5}

	assert(linearSearch(arr, 5), true, "linear search found failed")
	assert(linearSearch(arr, 10), false, "linear search not found failed")
}

func linearSearch(arr []int, v int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == v {
			return true
		}
	}
	return false
}

func assert[T comparable](a T, b T, msg string) {
	if a != b {
		panic(fmt.Sprintf("Assert Error: %s", msg))
	}
}
