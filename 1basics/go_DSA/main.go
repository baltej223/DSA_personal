package main

import "gods/linkedlist"

func main() {
	var list *linkedlist.Node = linkedlist.Create(7, 4279)
	linkedlist.Print(list)
}
