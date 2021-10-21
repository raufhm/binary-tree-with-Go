package main

import (
	"fmt"
	"strconv"
)

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
	parentChildData := []int{90, 10, 200, 300, 350, 40, 34, 67, 121, 69}
	for _, data := range parentChildData {
		tree.Insert(data)
	}
	line := "==========================================================================="
	fmt.Println(line)
	listData := fmt.Sprintf("this is the root data [%v],\nthis is parent and child data %v", tree.Key, parentChildData)
	fmt.Println(listData)

	// search the data
	searchData := []int{75, 69, 90}
	fmt.Println(line)

	for _, data := range searchData {
		found := tree.Search(data)
		strData := strconv.Itoa(data)
		if !found {
			result := fmt.Sprintf("is %s exist ? No, %s is not exist 😭", strData, strData)
			fmt.Println(result)
		} else {
			result := fmt.Sprintf("is %s exist ? Yes, %s is exist 😍", strData, strData)
			fmt.Println(result)
		}
	}
}
