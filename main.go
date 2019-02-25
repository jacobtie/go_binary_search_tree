package main

import(
	"fmt"
)

func main() {
	runMenu()
}

func runMenu() {
	var num int
	bst := &BinarySearchTree{}

	for num != 8 {
		fmt.Println("\nChoose an option from the following list")
		fmt.Println("1. Insert Numbers")
		fmt.Println("2. Run In Order Traveral")
		fmt.Println("3. Run Pre Order Traveral")
		fmt.Println("4. Run Post Order Traveral")
		fmt.Println("5. Get Max")
		fmt.Println("6. Get Min")
		fmt.Println("7. Contains")
		fmt.Println("8. Quit\n")
		
		fmt.Scanf("%d", &num)

		switch num {
		case 1:
			// bst.insert(getNum())
			bst.insert(insertNums()...)
		case 2:
			fmt.Println(bst.inOrderTraversal())
		case 3:
			fmt.Println(bst.preOrderTraversal())
		case 4:
			fmt.Println(bst.postOrderTraversal())
		case 5:
			fmt.Println(bst.getMax())
		case 6:
			fmt.Println(bst.getMin())
		case 7:
			search(bst)
		}
	}
}

func insertNums() []int {
	num := -1
	var arr []int

	for num != 0 {
		fmt.Println("Enter 1 to enter a number. Enter 0 to stop entering numbers.")
		fmt.Scanf("%d", &num)
		if num == 1 {
			arr = append(arr, getNum())
		}
	}

	return arr
}

func getNum() int {
	var num int

	fmt.Print("Please enter an integer: ")
	fmt.Scanf("%d", &num)
	
	return num
}

func search(bst *BinarySearchTree) {
	if bst.contains(getNum()) {
		fmt.Println("The number was found")
	} else {
		fmt.Println("The number was not found")
	}
}
