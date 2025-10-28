package create

import "trees/definitions"

type Node = definitions.Node

func Create(arr []int) Node {
	var root Node
	root.Data = arr[0]

	for i := 1; i < len(arr); i++ {
		var current_level *Node
		current_level = &root
		for {
			if arr[i] < current_level.Data {
				if current_level.Left != nil {
					current_level = current_level.Left
				} else {
					current_level.Left = &Node{Data: arr[i], Left: nil, Right: nil}
					break
				}
			} else if arr[i] > current_level.Data {
				if current_level.Right != nil {
					current_level = current_level.Right
				} else {
					current_level.Right = &Node{Data: arr[i], Left: nil, Right: nil}
					break
				}
			} else {
				// Duplicates are ignored!
				break
			}
		}
	}

	return root
}
