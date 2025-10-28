package traversal

import (
	"fmt"
	"trees/definitions"
)

// PLR

type Node = definitions.Node

func PreOrder(root *Node) {
	// Okay two ways to do it. Either to use stack for explicit iterative approach
	// Or use the recursion!

	if root == nil {
		return
	}
	fmt.Print(root.Data, " ") // Visit the node
	PreOrder(root.Left)       // Traverse left
	PreOrder(root.Right)      // Traverse right
}
