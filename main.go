package main

import (
	"fmt"
	"godatastructures/linkedlist"
	"godatastructures/tree"
)

func main() {
	tests := map[string]bool{
		"test_linkedlist": false,
		"test_tree":       true,
	}

	if tests["test_linkedlist"] {
		linkedlist.LinkedListTest()
	} else {
		fmt.Println("[SKIP] LinkedList test")
	}
	if tests["test_tree"] {
		tree.BinarySearchTreeTest()
	} else {
		fmt.Println("[SKIP] BinarySearchTree test")
	}
}
