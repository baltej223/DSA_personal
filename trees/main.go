package main

import (
	"fmt"
	"trees/traversal"
	createTree "trees/trees"
)

func main() {

	arr := []int{60, 10, 90, 40, 80, 20, 70, 50, 100, 30}

	tr := createTree.Create(arr)
	traversal.PreOrder(&tr)
	fmt.Println("")
	traversal.InOrder(&tr)
	fmt.Println("")
	traversal.PostOrder(&tr)
	fmt.Println("")

}
