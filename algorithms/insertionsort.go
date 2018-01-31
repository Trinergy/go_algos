package main

import (
	"fmt"
)

// WIP
// take first element
// compare it with other elements in same list
// if it is bigger than element at [i], insert to right, else left

// InsertionSort takes a list of integers and sorts it in ascending order
func InsertionSort(list []int) []int {
	sortedList := make([]int, len(list))
	sortedList[0], list = list[0], list[1:]
	fmt.Printf("\n sortedList222: %v \n list222: %v", sortedList, list)
	for i, value := range list {
		if value < sortedList[i] {
			sortedList = insert(sortedList, i-1, value)
		} else {
			sortedList = insert(sortedList, i+1, value)
		}
	}
	return sortedList
}

func insert(list []int, index int, value int) []int {
	if index > len(list) {
		fmt.Println("Index out of range")
	}
	insertedList := make([]int, len(list)+1)

	copy(insertedList[:index], list[:index])
	insertedList[index] = value
	copy(insertedList[index+1:], list[index:])

	return insertedList
}

func main() {
	list := []int{2, 6, 1, 3, 5, 8, 10, 23}
	fmt.Printf("\n unsorted: %v", list)
	sortedList := InsertionSort(list)
	fmt.Printf("\n sorted: %v", sortedList)

	// Test insert into array
	// testList := []int{1, 2, 4}
	// fmt.Printf("\n Test List: %v", testList)
	// idx := 2
	// value := 3
	// insertedList := insert(testList, idx, value)
	// fmt.Printf("\n inserted list: %v", insertedList)
}
