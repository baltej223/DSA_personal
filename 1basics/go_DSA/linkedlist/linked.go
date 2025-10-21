package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

func Create(size int, def_val int) *Node {
	var prev *Node
	var first = &Node{def_val, nil}
	prev = first
	for i := 0; i < size-1; i++ {
		nthNode := Node{def_val, nil}
		prev.next = &nthNode
		prev = &nthNode
	}
	return first
}

func CreateFromSlice(arr []int) *Node {
	var length int = len(arr)
	var prev *Node
	var first = &Node{arr[0], nil}
	prev = first
	for i := 1; i < length; i++ {
		nthNode := &Node{arr[i], nil}
		prev.next = nthNode
		prev = nthNode
	}
	return first
}

func Print(list *Node) {
	// fmt.Println("Print called")
	current := list
	// fmt.Println(current)
	for next := current.next; next != nil; next = next.next {
		fmt.Println(current.data)
	}

}
