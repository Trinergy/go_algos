package main

import (
	"fmt"
)

type DLL struct {
	value    string
	previous *DLL
	next     *DLL
}

func (dll *DLL) add(value string) {
	nextDLL := DLL{value: value, previous: dll}
	dll.next = &nextDLL
}

func (dll *DLL) remove() {
	previousDLL := dll.previous
	nextDLL := dll.next

	if previousDLL.previous == dll { // removing from head
		previousDLL.previous = nil
	}

	previousDLL.next = nextDLL
}

func main() {
	headDLL := DLL{value: "chicken"}
	headDLL.add("beef")
	nextDLL := headDLL.next
	headDLL.previous = nextDLL

	fmt.Printf("verify add: \nthis is first value: %q", headDLL.value)
	fmt.Printf("\nthis is the second value: %q", nextDLL)
	fmt.Printf("\nthis is the previous DLL: %q", headDLL.previous)

	nextDLL.remove()

	fmt.Printf("\nverify remove:")
	fmt.Printf("\n this is the DLL after removal: %q", headDLL)
}
