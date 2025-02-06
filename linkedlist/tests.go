package linkedlist

import (
	"fmt"
	"math/rand"
	"time"
)

func LinkedListTest() {
	fmt.Println("[START] LinkedList test")
	// we will compare the performance of SingleLinkedList and DoubleLinkedList
	var SingleLinkedList = SingleLinkedList[int]{}
	var DoubleLinkedList = DoubleLinkedList[int]{}
	appendCount := 1_000_000
	getCount := 5_000
	start := time.Now()
	for counter := 0; counter < appendCount; counter++ {
		SingleLinkedList.Append(counter)
	}
	fmt.Printf("SingleLinkedList Append %d item took %s\n", appendCount, time.Since(start))

	start = time.Now()
	for counter := 0; counter < appendCount; counter++ {
		DoubleLinkedList.Append(counter)
	}
	fmt.Printf("DoubleLinkedList Append %d item took %s\n", appendCount, time.Since(start))
	start = time.Now()
	for counter := 0; counter < getCount; counter++ {
		randomIndex := rand.Intn(appendCount)
		SingleLinkedList.ItemAt(randomIndex)
	}
	fmt.Printf("SingleLinkedList Get %d random item took %s\n", getCount, time.Since(start))

	start = time.Now()
	for counter := 0; counter < getCount; counter++ {
		randomIndex := rand.Intn(appendCount)
		DoubleLinkedList.ItemAt(uint(randomIndex))
	}
	fmt.Printf("DoubleLinkedList Get %d random item took %s\n", getCount, time.Since(start))
	fmt.Println("[END] LinkedList test")
	// DoubleLinkedList Get will be faster than SingleLinkedList Get because of the way we implemented ItemAt method
	// DoubleLinkedList can search from the end of the list if the index is closer to the end
	// If total size grows, the difference will be more significant
	// because SingleLinkedList will have to search from the beginning every time
	// and the gap between the index and the end of the list will be bigger

	// now we will check memory usage
	// each single linked list has a head and a tail pointer with size 8 bytes each = 16 bytes
	// each single node has a data with size X and a next pointer with size 8 bytes = X + 8 bytes

	// each double linked list has a head, a tail pointer and a size with size 8 bytes each = 24 bytes
	// each double linked node has a data with size X and a next, prev pointer with size 8 bytes each = X + 16 bytes

	// for 100_000_000 items of int64 type (8 bytes) we will have
	// SingleLinkedList: 16 + 1_000_000 * (8 + 8) = 16 + 16_000_000 = 16_000_016 bytes = 16 MB
	// DoubleLinkedList: 24 + 1_000_000 * (8 + 16) = 24 + 24_000_000 = 24_000_024 bytes = 24 MB
	// 24/16 = 1.5 times more memory usage for DoubleLinkedList

	//get 5000 random item in list with 1_000_000 item took:
	// 3.35 sec for SingleLinkedList
	// 2.29 sec for DoubleLinkedList
	// 3.35/2.29 = 1.46 times faster for DoubleLinkedList
	// Note : if we increase the number of items in the list, the difference will change
	// and the DoubleLinkedList will be faster more than 1.46 times
	// but the memory usage will be exactly 1.5 times more than SingleLinkedList
	// so for a large number of items, we can use DoubleLinkedList if we need to get items faster
	// but for a small number of items, we can use SingleLinkedList if we need to save memory
}
