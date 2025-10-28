package linkedlist

import "fmt"

type DLLNode struct {
	Data int
	Next *DLLNode
	Prev *DLLNode
}

func CreateDoublyLinkedListFromSlice(arr []int) *DLLNode {
	var length int = len(arr)
	var prev *DLLNode
	var first = &DLLNode{arr[0], nil, nil}
	prev = first
	for i := 1; i < length; i++ {
		nthDLLNode := &DLLNode{arr[i], nil, nil}
		prev.Next = nthDLLNode
		prev = nthDLLNode
		nthDLLNode.Prev = prev
	}
	return first
}

func PrintDLL(dllNode *DLLNode) {
	for current := dllNode; current != nil; current = current.Next {
		fmt.Print(current.Data, " ")
	}
}
