package main

import (
	"fmt"
)

func main() {
	var items = []int{1, 4, 7, 9, 2, 5, 3}
	qS := quickSort(items)
	fmt.Println("qS:", qS)

	low := 0
	high := len(qS) - 1
	target := 1
	bS := binarySearch(qS, target, low, high)
	fmt.Println("bS:", bS)

	// interface
	cir := circle{radius: 3}
	calculate(cir)
	rec := rect{width: 4, height: 6}
	calculate(rec)

	// node
	root := &Node{data: 8}
	root.insert(12)
	root.insert(4)
	root.insert(9)
	root.insert(2)
	root.insert(5)
	root.printNodes()

	fmt.Println("Inorder traversal:")
	root.inOrderTraversal()
	fmt.Println("\nSearch for 5:")
	result := root.search(5)
	if result != nil {
		fmt.Println("Found:", result.data)
	} else {
		fmt.Println("Not found.")
	}
}
