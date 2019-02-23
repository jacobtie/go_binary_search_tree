package main 

type BinarySearchTree struct {
	root *BinaryTreeNode
	numItems int
}

func (bst *BinarySearchTree) insert(elems ...int) {
	for _, v := range elems {
		node := &BinaryTreeNode{nil, nil, v}
		if bst.root == nil {
			bst.root = node
		} else {
			bst.insert_helper(bst.root, node)
		}
		bst.numItems += 1
	}
}

func (bst *BinarySearchTree) insert_helper(subRoot *BinaryTreeNode, node *BinaryTreeNode) {
	if node.Element <= subRoot.Element {
		if subRoot.Left == nil {
			subRoot.Left = node
		} else {
			bst.insert_helper(subRoot.Left, node)
		}
	} else {
		if subRoot.Right == nil {
			subRoot.Right = node
		} else {
			bst.insert_helper(subRoot.Right, node)
		}
	}
}

func (bst *BinarySearchTree) inOrderTraversal() []int {
	results := make([]int, 0, bst.numItems)
	bst.inOrderTraversal_helper(bst.root, &results)
	return results
}

func (bst *BinarySearchTree) inOrderTraversal_helper(subRoot *BinaryTreeNode, results *[]int) {
	if subRoot == nil {
		return
	}
	bst.inOrderTraversal_helper(subRoot.Left, results)
	*results = append(*results, subRoot.Element)
	bst.inOrderTraversal_helper(subRoot.Right, results)
}
