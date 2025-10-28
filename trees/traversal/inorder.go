package traversal

import (
	"fmt"
	"trees/definitions"
)

// LPR
type Tree = *definitions.Tree

func InOrder(root Tree) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Print(root.Data, " ")
	InOrder(root.Right)
}
