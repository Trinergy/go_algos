package main

import (
	"fmt"
)

/* TODO: Make the List manage its items. Not have the items know of their parent
* Rename for clarity
 */
type ArrayDLL struct {
	length int
}

type DLL struct {
	value    string
	previous *DLL
	next     *DLL
	parent   *ArrayDLL
}

func (dll *DLL) add(value string) {
	parentDLL := dll.parent
	nextDLL := DLL{value: value, previous: dll, parent: parentDLL}
	dll.next = &nextDLL

	parentDLL.length++
}

func (dll *DLL) remove() {
	previousDLL := dll.previous
	nextDLL := dll.next

	parentDLL := dll.parent
	parentDLL.length--
	parentDLL = nil

	if previousDLL.previous == dll { // removing from head
		previousDLL.previous = nil
	}

	previousDLL.next = nextDLL
}

func main() {
	arrayDLL := ArrayDLL{length: 1} // assumes to be initiated with a child (bad)
	// currently child also assumes knowledge of parent
	headDLL := DLL{value: "chicken", parent: &arrayDLL}
	headDLL.add("beef")
	nextDLL := headDLL.next
	headDLL.previous = nextDLL

	fmt.Printf("verify add: \nthis is first value: %q", headDLL.value)
	fmt.Printf("\nthis is the second value: %q", nextDLL)
	fmt.Printf("\nthis is the previous DLL: %q", headDLL.previous)
	fmt.Printf("\nthis is the arrayDLL's length: %d", arrayDLL.length)

	nextDLL.remove()

	fmt.Printf("\nverify remove:")
	fmt.Printf("\n this is the DLL after removal: %q", headDLL)
	fmt.Printf("\nthis is the arrayDLL's length: %d", arrayDLL.length)
}
