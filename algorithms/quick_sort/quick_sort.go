package main

import (
	"fmt"
	"math/rand"
)

// QuickSort sorts the array's elements in ascending order from lowest to highest
func QuickSort(list []int) []int {
	length := len(list)
	// base case
	if length <= 1 {
		return list
	}

	// recursive case
	pivot := list[rand.Intn(length)]
	var leftList, rightList []int

	for _, v := range list {
		if v <= pivot {
			leftList = append(leftList, v)
		} else {
			rightList = append(rightList, v)
		}
	}

	return append(QuickSort(leftList), QuickSort(rightList)...)
}

func main() {
	sampleData := []int{1, 88, 99, 433, 20, 44, 3345, 4354352, 45, 2}
	fmt.Printf("\n this is the unsorted list: %v", sampleData)
	sortedData := QuickSort(sampleData)
	fmt.Printf("\n this is the sorted list: %v", sortedData)
}
