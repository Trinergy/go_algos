package main

import "fmt"

// QuickSort sorts the array's elements in ascending order from lowest to highest
func QuickSort(list []int) []int {
	length := len(list)
	// base case
	if length <= 1 {
		return list
	}

	// recursive case
	pivot := list[(length / 2)]
	var leftList, rightList []int
	for _, v := range list {
		if v <= pivot {
			leftList = append(leftList, v)
		} else {
			rightList = append(rightList, v)
		}
	}

	return merge(QuickSort(leftList), QuickSort(rightList))
}

func merge(list1, list2 []int) []int {
	return append(list1, list2...)
}

func main() {
	sampleData := []int{1, 88, 99, 433, 20, 44, 3345, 4354352, 45, 2}
	fmt.Printf("\n this is the unsorted list: %v", sampleData)
	sortedData := QuickSort(sampleData)
	fmt.Printf("\n this is the sorted list: %v", sortedData)
}
