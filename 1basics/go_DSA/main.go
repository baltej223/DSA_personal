package main

import (
	"fmt"
	"go_ds/linkedlist"
	"go_ds/search"
)

func main() {
	var list *linkedlist.Node = linkedlist.Create(7, 4279)
	linkedlist.Print(list)

	slice := []int{1, 2, 425, 646, 43, 3, 4215, 643, 67, 42352, 632, 634, 7}
	var list2 *linkedlist.Node = linkedlist.CreateFromSlice(slice)
	linkedlist.Print(list2)

	slice = []int{1, 2, 3, 4, 5, 6, 7, 10}
	ll := linkedlist.CreateFromSlice(slice)
	fmt.Println("Found element at index:", search.LinearSearch(ll, 10))

	slice = []int{12, 42, 53, 62, 52, 51, 74, 75}
	dll := linkedlist.CreateDoublyLinkedListFromSlice(slice)
	linkedlist.PrintDLL(dll)

	// search.BinarySearch(dll, 42)
	// var sorted_slice []int = []int{1, 3, 5, 6, 7, 13, 52, 56, 89, 141}
	// fmt.Println(linkedlist.CountLength(dll))
	// dll = linkedlist.CreateDoublyLinkedListFromSlice(sorted_slice)
	// var pos int = search.BinarySearch(dll, 42)
	// fmt.Println(pos)
}
