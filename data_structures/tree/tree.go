package main

import (
	"fmt"
	"os"

	"github.com/Trinergy/gologger"
)

var (
	log    = gologger.SetupLogFile("debug_log.txt")
	logger = gologger.SetupLogger(log)
)

// Tree represents a collection of nodes
type Tree interface {
	size() int
	isRoot(n Node) bool
	isLeafNode(n Node) bool
	swap(one Node, two Node)
	replace(old Node, new Node)
	sort()
}

// BinaryTree is the collection of 2-set nodes
type BinaryTree struct {
	root *Node
}

// Node is one point of data in a tree
type Node struct {
	id       int
	value    int
	parent   *Node // TODO: No need for parent if purpose is only to know if it is a root. Resolved at tree level
	children []Node
}

func NewBinaryTree(list []int) BinaryTree {
	if len(list) == 0 {
		fmt.Printf("list can not be empty: %v", list)
	}
	nodeList := convertToNodes(list)
	childIndex := 1
	for _, node := range nodeList {
		for i := 0; i < 2; i++ {
			if childIndex == len(nodeList)-1 {
				break
			}
			node.children[i] = nodeList[childIndex]
			childIndex++
		}
	}
	return BinaryTree{&nodeList[0]}
}

func convertToNodes(list []int) []Node {
	nodeList := make([]Node, len(list))
	for i, v := range list {
		nodeList[i] = Node{id: i, value: v, children: make([]Node, 2)}
	}
	return nodeList
}

// NewBinaryTree2 constructs a binary tree given a list of integers
func NewBinaryTree2(list []int) BinaryTree {
	if len(list) == 0 {
		fmt.Printf("list can not be empty: %v", list)
		os.Exit(1)
	}

	root := Node{id: 0, value: list[0], children: make([]Node, 2)}
	queue := []Node{root} // list of potential foster parents who can adopt child nodes

	const limit = 2
	count := 0
	for i, v := range list[1:] {
		if count == limit {
			logger.Printf("%+v", queue[0])
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
	defer log.Close()

	sample := []int{8, 5, 45, 66, 78, 7587, 56, 20}
	bt := NewBinaryTree(sample)
	// fmt.Println(bt.root.children[1].children)
	fmt.Println(bt.root.children)
}
