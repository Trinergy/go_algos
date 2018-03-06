package main

import (
	"fmt"
)

/* TODO: Make the List manage its items. Not have the items know of their parent
* Rename for clarity
 */
type DLList struct {
	length    int
	firstItem *DLLItem
	lastItem  *DLLItem
}

type DLLItem struct {
	index    int
	value    string
	previous *DLLItem
	next     *DLLItem
}

func (dll *DLList) add(value string) {
	// New List
	if dll.lastItem == nil {
		firstItem := DLLItem{index: 0, value: value}
		dll.firstItem = &firstItem
		dll.lastItem = &firstItem
	} else {
		lastItem := dll.lastItem
		newItem := DLLItem{index: lastItem.index + 1, value: value, previous: lastItem, next: dll.firstItem}
		lastItem.next = &newItem
		dll.lastItem = &newItem
	}

	dll.length++
}

func (dll *DLList) removeByIndex(index int) *DLLItem {
	if index > dll.lastItem.index {
		panic("Index out of range")
	}

	if index == 0 {
		// handle removing first item
		return dll.firstItem
	}
	if index == dll.lastItem.index {
		removeItem := dll.lastItem
		dll.lastItem = removeItem.previous
		dll.removeItem(removeItem)
		return removeItem
	}

	var i = 0
	var removedItem *DLLItem
	for i < dll.length {
		nextItem := dll.firstItem.next
		if nextItem.index == index {
			removedItem = nextItem
			dll.removeItem(nextItem)
			dll.reindexFrom(nextItem)
			dll.length--
			break
		}
		i++
	}

	return removedItem
}

func (dll *DLList) removeItem(item *DLLItem) {
	nextItem := item.next
	previousItem := item.previous

	nextItem.previous = previousItem
	previousItem.next = nextItem
	// fmt.Printf("\nthis is the nextItem's index: %d", nextItem.index)
}

func (dll *DLList) reindexFrom(startItem *DLLItem) {
	// don't reindex if from start of list
	if startItem.index == 0 {
		return
	}

	var i = 0
	var max = dll.length - startItem.index
	// fmt.Printf("\nthis is the max: %d", max)
	// fmt.Printf("\nthis is the startItem's index: %d", startItem.index)

	startItem.index--
	for i < max {
		nextItem := startItem.next
		nextItem.index--
		i++
	}
}

func main() {
	list := DLList{}

	fmt.Println("\n=== ADD ===")
	list.add("chicken")
	fmt.Printf("\nthis is the last DLLItem: %q", list.lastItem.value)
	fmt.Printf("\nthis is the last DLLItem's index: %d\n", list.lastItem.index)
	fmt.Printf("\nthis is the list's length: %d\n", list.length)

	list.add("beef")
	fmt.Printf("\nthis is the last DLLItem: %q", list.lastItem.value)
	fmt.Printf("\nthis is the last DLLItem's index: %d\n", list.lastItem.index)
	fmt.Printf("\nthis is the list's length: %d\n", list.length)

	list.removeByIndex(1)
	fmt.Println("\n=== REMOVE BY INDEX ===")
	fmt.Printf("\nthis is the last DLLItem: %q", list.lastItem.value)
	fmt.Printf("\nthis is the last DLLItem's index: %d\n", list.lastItem.index)
	fmt.Printf("\nthis is the list's length: %d\n", list.length)

	fmt.Println("\n=== ADD BACK ===")
	list.add("pork")
	fmt.Printf("\nthis is the last DLLItem: %q", list.lastItem.value)
	fmt.Printf("\nthis is the last DLLItem's index: %d\n", list.lastItem.index)
	fmt.Printf("\nthis is the list's length: %d\n", list.length)

	list.add("lamb")
	fmt.Printf("\nthis is the last DLLItem: %q", list.lastItem.value)
	fmt.Printf("\nthis is the last DLLItem's index: %d\n", list.lastItem.index)
	fmt.Printf("\nthis is the list's length: %d\n", list.length)

	list.removeByIndex(1)
	fmt.Println("\n=== REMOVE BY INDEX ===")
	fmt.Printf("\nthis is the last DLLItem: %q", list.lastItem.value)
	fmt.Printf("\nthis is the list's length: %d\n", list.length)
}
