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

	for num != 5 {
		fmt.Println("\nChoose an option from the following list")
		fmt.Println("1. Insert Numbers")
		fmt.Println("2. Run In Order Traveral")
		fmt.Println("3. Run Pre Order Traveral")
		fmt.Println("4. Run Post Order Traveral")
		fmt.Println("5. Quit\n")
		
		fmt.Scanf("%d", &num)

		fmt.Print("\n")

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
