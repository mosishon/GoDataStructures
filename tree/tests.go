package tree

import (
	"fmt"
	"math/rand"
	"time"
)

func BinarySearchTreeTest() {
	fmt.Println("[START] BinarySearchTree test")
	var bst = &BinarySearchTree[int]{}
	comparator := func(a, b int) ComparatorResult {
		switch {
		case a < b:
			return LessThan
		case a > b:
			return Greater
		default:
			return Equal
		}
	}

	appendCount := 100_000
	getCount := 5_000

	start := time.Now()
	for counter := 0; counter < appendCount; counter++ {
		bst.Insert(counter, comparator)
	}
	fmt.Printf("BinarySearchTree Insert %d items took %s\n", appendCount, time.Since(start))

	start = time.Now()
	for counter := 0; counter < getCount; counter++ {
		randomIndex := rand.Intn(appendCount)
		bst.Find(randomIndex, comparator)
	}
	fmt.Printf("BinarySearchTree Find %d random items took %s\n", getCount, time.Since(start))
	fmt.Println("[END] BinarySearchTree test")

	// now we will check memory usage
	// each node in the binary search tree has a data with size X and left, right pointers with size 8 bytes each = X + 16 bytes
	// the binary search tree itself has a root pointer with size 8 bytes

	// for 1_000_000 items of int64 type (8 bytes) we will have
	// BinarySearchTree: 8 + 1_000_000 * (8 + 16) = 8 + 24_000_000 = 24_000_008 bytes = 24 MB

	// --*-----*----------*-----------*------------*----------*----------*---------*

	// BinarySearchTreeTest performs a test on the BinarySearchTree implementation.
	// It inserts a large number of items into the tree and then performs random
	// searches to measure the performance of these operations.
	//
	// The time complexity of the Insert operation in a Binary Search Tree (BST) is
	// O(log n) on average, but it can degrade to O(n) in the worst case if the tree
	// becomes unbalanced.
	//
	// The time complexity of the Find operation in a Binary Search Tree (BST) is
	// O(log n) on average, but it can degrade to O(n) in the worst case if the tree
	// becomes unbalanced.
	//
	// The test inserts 1,000,000 items into the BST and then performs 5,000 random
	// searches, measuring the time taken for each operation.
	// --*-----*----------*-----------*------------*----------*----------*---------*

}
