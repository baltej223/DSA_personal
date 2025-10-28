package traversal

import (
	"fmt"
)

// LRP
func PostOrder(root Tree) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Print(root.Data, " ")
}
