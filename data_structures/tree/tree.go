package main

import "fmt"

// Tree represents a collection of nodes
type Tree interface {
	size() int
	isRoot(n Node) bool
	isLeafNode(n Node) bool
}

// BinaryTree is the collection of 2-set nodes
type BinaryTree struct {
	Tree
	root *Node
}

// Node is one point of data in a tree
type Node struct {
	id       int
	value    int
	parent   *Node
	children []Node
}

// NewBinaryTree constructs a binary tree given a list of integers
func NewBinaryTree(list []int) BinaryTree {
	if len(list) <= 0 {
		fmt.Printf("list can not be empty: %v", list)
	}

	root := Node{id: 0, value: list[0], children: make([]Node, 2)}
	queue := []Node{root} // list of potential foster parents who can adopt child nodes

	const limit = 2
	count := 0
	for i, v := range list[1:] {
		if count == limit {
			queue = queue[1:]
			count = 0
		}

		parent := &queue[0]
		node := Node{id: i, value: v, parent: parent, children: make([]Node, 2)}
		parent.children[count] = node

		queue = append(queue, node)
		count++
	}

	return BinaryTree{root: &root}
}

func main() {
	sample := []int{8, 5, 45, 66, 78, 7587, 56}
	bt := NewBinaryTree(sample)
	fmt.Println(bt.root.children[0].children)
	fmt.Println(bt.root)
}
