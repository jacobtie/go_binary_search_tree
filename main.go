package main

import "fmt"

func main() {
	arr := []int{6, 4, 65, 2, 9, 1, 8}
	bst := &BinarySearchTree{nil, 0}
	bst.insert(arr...)
	fmt.Println(bst)
	fmt.Println(bst.inOrderTraversal())
}
