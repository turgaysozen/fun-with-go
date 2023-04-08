package main

func binarySearch(items []int, target, low, high int) int {
	if low > high {
		return -1
	}

	midIdx := (low + high) / 2

	if items[midIdx] == target {
		return midIdx
	} else if items[midIdx] < target {
		return binarySearch(items, target, midIdx+1, high)
	} else {
		return binarySearch(items, target, low, midIdx-1)
	}
}
