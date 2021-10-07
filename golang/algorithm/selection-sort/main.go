package main

import "fmt"

func main() {
	arr := []int{5, 3, 2, 1, 4}
	newArr := selectionSort(arr)
	fmt.Println(newArr)
}

func findSmallestIndex(arr []int) int {
	smallestIndex := 0
	smallest := arr[smallestIndex]

	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallestIndex = i
			smallest = arr[i]
		}
	}

	return smallestIndex
}

func selectionSort(arr []int) []int {
	newArr := make([]int, len(arr))
	for i := 0; i < len(newArr); i++ {
		smallestIndex := findSmallestIndex(arr)
		newArr[i] = arr[smallestIndex]
		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
	}
	return newArr
}
