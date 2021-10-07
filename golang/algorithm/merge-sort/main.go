package main

import "fmt"

func main() {
	arr := []int{8, 4, 5, 7, 1, 3, 6, 2}
	mergeSort(arr)
	fmt.Println(arr)
}

func mergeSort(arr []int) {
	// 统一创建临时空间
	temp := make([]int, len(arr))
	sort(arr, 0, len(arr)-1, temp)
}

func sort(arr []int, left, right int, temp []int) {
	if left < right {
		mid := (left + right) / 2
		sort(arr, left, mid, temp)
		sort(arr, mid+1, right, temp)
		merge(arr, left, mid, right, temp)
	}
}

func merge(arr []int, left, mid, right int, temp []int) {
	copy(temp[left:right+1], arr[left:right+1])

	i := left
	j := mid + 1
	for t := left; t <= right; t++ {
		if i > mid {
			// 左边无数据可比较
			// 将右边数据依次放入
			arr[t] = temp[j]
			j++
		} else if j > right {
			// 右边无数据比较
			// 将左边数据依次放入
			arr[t] = temp[i]
			i++
		} else {
			// 左右两边都有数据
			// 比较先保存较小的数据
			if temp[i] <= temp[j] {
				arr[t] = temp[i]
				i++
			} else {
				arr[t] = temp[j]
				j++
			}
		}

	}
}
