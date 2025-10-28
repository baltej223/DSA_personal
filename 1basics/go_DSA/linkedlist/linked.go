package linkedlist

import (
	"fmt"
	// "go_ds/linkedlist"
)

type Node struct {
	Data int
	Next *Node
}

func Create(size int, def_val int) *Node {
	var prev *Node
	var first = &Node{def_val, nil}
	prev = first
	for i := 0; i < size-1; i++ {
		nthNode := &Node{def_val, nil}
		prev.Next = nthNode
		prev = nthNode
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
		prev.Next = nthNode
		prev = nthNode
	}
	return first
}

func Print(list *Node) {
	for current := list; current != nil; current = current.Next {
		fmt.Println(current.Data)
	}
}

func CountLength(list interface{}) int {
	count := 0
	switch l := list.(type) {
	case *Node:
		for current := l; current != nil; current = current.Next {
			count++
		}
	case *DLLNode:
		for current := l; current != nil; current = current.Next {
			count++
		}
	}
	return count
}
