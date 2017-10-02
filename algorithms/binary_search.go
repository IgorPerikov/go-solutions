package main

import (
	"math"
	"fmt"
)

func main() {
	fmt.Println(initSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10), 9)
	fmt.Println(initSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9), 8)
	fmt.Println(initSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 105), -1)
	fmt.Println(initSearch([]int{1}, 1), 0)
	fmt.Println(initSearch([]int{0, 1}, 1), 1)
	fmt.Println(initSearch([]int{1}, 2), -1)
	fmt.Println(initSearch([]int{1, 2}, 3), -1)
	fmt.Println(initSearch([]int{1, 3}, 2), -1)
	fmt.Println(initSearch([]int{1, 3, 5, 100}, 2), -1)
	fmt.Println(initSearch([]int{1, 3, 5, 100}, 5), 2)
	fmt.Println(initSearch([]int{1, 3, 5, 100}, 100), 3)
	fmt.Println(initSearch([]int{1, 3, 5, 100}, 1), 0)
}

func initSearch(array []int, search int) int {
	if len(array) == 0 {
		return -1
	}
	return binarySearch(array, search, 0, len(array)-1)
}

func binarySearch(array []int, search, low, high int) int {
	if high == low && array[high] != search {
		return -1
	}

	pivotIndex := int(math.Floor(float64(high-low)/2)) + low
	switch {
	case array[pivotIndex] == search:
		return pivotIndex
	case search > array[pivotIndex]:
		return binarySearch(array, search, pivotIndex+1, high)
	default: // case array[pivotIndex] > search:
		return binarySearch(array, search, low, pivotIndex)
	}
}
