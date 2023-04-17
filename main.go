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
}
