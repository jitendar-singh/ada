package main

import "fmt"

func binarySearch(dataList []int, key int) bool {
	low := 0
	high := len(dataList) - 1

	for low <= high {
		median := (low + high) / 2
		fmt.Println("New median is %s", median)
		if dataList[median] < key {
			low = median + 1
			fmt.Println("New low is %s", low)
		} else {
			high = median - 1
			fmt.Println("New high is %s", high)
		}
	}

	if low == len(dataList) || dataList[low] != key {
		return false
	}
	return true
}

func main() {
	items := []int{1, 2, 5, 7, 9, 23, 45, 67, 77, 88, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}
	fmt.Print(binarySearch(items, 94))
}
