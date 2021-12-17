package main

import "fmt"

func ShellSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}

	gap := n / 2
	for gap > 0 {
		for i := gap; i < n; i++ {
			for j := i; j >= gap && arr[j] < arr[j-gap]; j = j - gap {
				arr[j-gap], arr[j] = arr[j], arr[j-gap]
			}
		}
		gap = gap / 2
	}
}

func main() {
	arr := []int{9, 1, 4, 6, 2, 3, 7}
	ShellSort(arr)
	fmt.Println(arr)
}
