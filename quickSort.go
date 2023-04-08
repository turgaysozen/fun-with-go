package main

func quickSort(items []int) []int {
	if len(items) <= 1 {
		return items
	}

	midIdx := len(items) / 2
	pivot := items[midIdx]
	var left, mid, right []int
	for _, i := range items {
		if pivot < i {
			left = append(left, i)
		} else if i == pivot {
			mid = append(mid, i)
		} else {
			right = append(right, i)
		}
	}

	left = quickSort(left)
	right = quickSort(right)
	sortedArr := append(append(right, mid...), left...) // sorting asc

	return sortedArr
}
