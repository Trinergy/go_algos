package main

import "fmt"

// MergeSort sorts the array's elements in ascending order from lowest to highest
func MergeSort(list []int) []int {
	length := len(list)
	// base case
	if length <= 1 {
		return list
	}

	// recursive case
	pivot := (length / 2)
	leftList := list[:pivot]
	rightList := list[pivot:]
	return merge(MergeSort(leftList), MergeSort(rightList))
}

func merge(list1, list2 []int) []int {
	maxLength := len(list1) + len(list2)
	orderedList := make([]int, maxLength)
	var j, k = 0, 0

	for i := 0; i < maxLength; i++ { // see if range works w/ 2 args
		if j > len(list1)-1 && k <= len(list2)-1 {
			orderedList[i] = list2[k]
			k++
		} else if k > len(list2)-1 && j <= len(list1)-1 {
			orderedList[i] = list1[j]
			j++
		} else if list1[j] < list2[k] {
			orderedList[i] = list1[j]
			j++
		} else {
			orderedList[i] = list2[k]
			k++
		}
	}

	return orderedList
}

func main() {
	sampleData := []int{1, 88, 99, 433, 20, 44, 3345, 4354352, 45, 2}
	fmt.Printf("\n this is the unsorted list: %v", sampleData)
	sortedData := MergeSort(sampleData)
	fmt.Printf("\n this is the sorted list: %v", sortedData)
}
