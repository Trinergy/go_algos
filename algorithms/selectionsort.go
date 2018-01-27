package main

import (
	"fmt"
)

// pseudocode
// find the lowest number in the list
// remove it from the list and add it to the sorted list
// repeat

// SelectionSort sorts the numbers from lowest to highest
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

func SecondSelectionSort(list []int) {
	size := len(list)
	for i := 0; i < size; i++ {
		minIdx := i
		for j := i + 1; j < size; j++ {
			if list[j] < list[minIdx] {
				minIdx = j
			}
		}
		list[i], list[minIdx] = list[minIdx], list[i]
	}
}

func main() {
	list := []int{3, 4, 5, 2, 1}
	var newList []int
	fmt.Printf("\n unsorted list %v", list)
	sortedList := SelectionSort(list, newList)
	fmt.Printf("\n sorted list: %v", sortedList)

	secondList := []int{10, 9, 8, 7, 6}
	fmt.Printf("\n secondList: %v", secondList)
	SecondSelectionSort(secondList)
	fmt.Printf("\n sorted secondList: %v", secondList)
}
