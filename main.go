package main

import "fmt"

func main() {
	arr := []int{28, 99, 42, 10, 11, 2, 53, 85}
	bst := &BinarySearchTree{nil, 0}
	bst.insert(arr...)
	fmt.Println(bst.inOrderTraversal())
	fmt.Println(bst.preOrderTraversal())
	fmt.Println(bst.postOrderTraversal())
}
