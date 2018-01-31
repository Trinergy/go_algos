package main

import (
	"fmt"
)

// SecondSelectionSort sorts the numbers from lowest to highest in the existing array (mutation)
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
	secondList := []int{10, 9, 8, 7, 6}
	fmt.Printf("\n secondList: %v", secondList)
	SecondSelectionSort(secondList)
	fmt.Printf("\n sorted secondList: %v", secondList)
}
