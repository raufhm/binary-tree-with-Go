package main

import "fmt"

// Component of Tree
type Node struct {
	Key   int
	Right *Node
	Left  *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right if input bigger than root value
		if n.Right == nil {
			n.Right = &Node{Key: k} // add value to right tree
		} else {
			n.Right.Insert(k) // if right data exist, then add new right tree data
		}
	} else if n.Key > k {
		// move left if input lower than root value
		if n.Left == nil {
			n.Left = &Node{Key: k} // add value to left tree
		} else {
			n.Left.Insert(k) // if left data exist, then add new left tree data
		}
	}
}

func (n *Node) Search(k int) bool {
	// check if Node is empty then directly return false as input not exist
	if n == nil {
		return false
	}
	if n.Key < k {
		// if input bigger than root, then call search function again to check next value
		return n.Right.Search(k)
	} else if n.Key > k {
		// if input lower than root, then call search function again to check next value
		return n.Left.Search(k)
	}
	// if data exist then return true
	return true
}

func main() {
	tree := &Node{Key: 100} // Initialize root value

	// insert parent and child value
	tree.Insert(120)
	tree.Insert(20)
	tree.Insert(84)
	tree.Insert(300)
	tree.Insert(623)
	tree.Insert(10)
	tree.Insert(70)
	tree.Insert(90)

	// search the data
	fmt.Println(tree.Search(700))
	fmt.Println(tree.Search(120))
	fmt.Println(tree.Search(100))
}
