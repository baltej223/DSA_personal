package search

import (
	"go_ds/linkedlist"
)

func LinearSearch(Head *linkedlist.Node, search_item int) int {
	position := -1
	couter := -1

	for current := Head; current != nil; current = current.Next {
		couter++
		if current.Data == search_item {
			position = couter
			break
		}
	}
	return position
}

// func main() {
// 	slice := []int{1, 2, 3, 4, 5, 6, 7, 10}
// 	ll := linkedlist.CreateFromSlice(slice)
// 	fmt.Println(LinearSearch(ll, 32))
// }
