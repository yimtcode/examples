package main

import "fmt"

func main() {
	arr := []int{3, 1, 9, 20, 4}
	binaryInsertSort(arr)
	fmt.Println(arr)
}

func binaryInsertSort(arr []int) {
	length := len(arr)
	var low, high, mid int

	for i := 1; i < length; i++ {
		t := arr[i]
		// binary search
		for low, high = 0, i-1; high >= low; {
			mid = (low + high) / 2
			if t < arr[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}

		// backward offset 1
		for j := i - 1; j >= low; j-- {
			arr[j+1] = arr[j]
		}
		// insert data
		arr[low] = t
	}
}

