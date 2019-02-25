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

func (bst *BinarySearchTree) traversal(t string) []int {
	results := make([]int, 0, bst.numItems)
	switch t {
	case "in":
		bst.inOrderTraversal_helper(bst.root, &results)
	case "pre":
		bst.preOrderTraversal_helper(bst.root, &results)
	case "post":
		bst.postOrderTraversal_helper(bst.root, &results)
	}
	return results
}

func (bst *BinarySearchTree) inOrderTraversal() []int {
	return bst.traversal("in")
}

func (bst *BinarySearchTree) preOrderTraversal() []int {
	return bst.traversal("pre")
}

func (bst *BinarySearchTree) postOrderTraversal() []int {
	return bst.traversal("post")
}

func (bst *BinarySearchTree) inOrderTraversal_helper(subRoot *BinaryTreeNode, results *[]int) {
	if subRoot == nil {
		return
	}
	bst.inOrderTraversal_helper(subRoot.Left, results)
	*results = append(*results, subRoot.Element)
	bst.inOrderTraversal_helper(subRoot.Right, results)
}

func (bst *BinarySearchTree) preOrderTraversal_helper(subRoot *BinaryTreeNode, results *[]int) {
	if subRoot == nil {
		return
	}
	*results = append(*results, subRoot.Element)
	bst.preOrderTraversal_helper(subRoot.Left, results)
	bst.preOrderTraversal_helper(subRoot.Right, results)
}

func (bst *BinarySearchTree) postOrderTraversal_helper(subRoot *BinaryTreeNode, results *[]int) {
	if subRoot == nil {
		return
	}
	bst.postOrderTraversal_helper(subRoot.Left, results)
	bst.postOrderTraversal_helper(subRoot.Right, results)
	*results = append(*results, subRoot.Element)
}

func (bst *BinarySearchTree) getMax() int {
	return bst.getMax_helper(bst.root)
}

func (bst *BinarySearchTree) getMax_helper(subRoot *BinaryTreeNode) int {
	if subRoot.Right == nil {
		return subRoot.Element
	}
	return bst.getMax_helper(subRoot.Right)
}

func (bst *BinarySearchTree) getMin() int {
	return bst.getMin_helper(bst.root)
}

func (bst *BinarySearchTree) getMin_helper(subRoot *BinaryTreeNode) int {
	if subRoot.Left == nil {
		return subRoot.Element
	}
	return bst.getMin_helper(subRoot.Left)
}

func (bst *BinarySearchTree) contains(num int) bool {
	return bst.contains_helper(bst.root, num)
}

func (bst *BinarySearchTree) contains_helper(subRoot *BinaryTreeNode, num int) bool {
	if subRoot == nil {
		return false
	} else if subRoot.Element == num {
		return true
	} else if subRoot.Element < num {
		return bst.contains_helper(subRoot.Right, num)
	} else {
		return bst.contains_helper(subRoot.Left, num)
	}
}
