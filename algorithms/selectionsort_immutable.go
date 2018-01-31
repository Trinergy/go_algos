package main

import (
	"fmt"
)

// pseudocode
// find the lowest number in the list
// remove it from the list and add it to the sorted list
// repeat

// SelectionSort sorts the numbers from lowest to highest in a new array (immutable)
func SelectionSort(list []int, sortedList []int) []int {
	// base case
	if len(list) == 0 {
		return sortedList
	}

	// recursive case
	lowest := list[0]
	var lowestIdx int
	tempList := make([]int, len(sortedList))
	copy(tempList, sortedList)

	// find the lowest number
	for i, value := range list {
		if value < lowest {
			lowest = value
			lowestIdx = i
		}
	}
	// add lowest number to new list
	sortedList = append(sortedList, list[lowestIdx])
	// remove lowest number from list
	list = append(list[:lowestIdx], list[lowestIdx+1:]...)

	//repeat
	return SelectionSort(list, sortedList)
}

func main() {
	list := []int{3, 4, 5, 2, 1}
	var newList []int
	fmt.Printf("\n unsorted list %v", list)
	sortedList := SelectionSort(list, newList)
	fmt.Printf("\n sorted list: %v", sortedList)
}
