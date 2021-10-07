package main

import "fmt"

func main() {
	arr := []int{24, 12, 34, 20, 53}
	straightInsertSort(arr)
	fmt.Println(arr)
}

// 直接插入排序
func straightInsertSort(arr []int) {
	length := len(arr)
	for i := 1; i < length; i++ {
		t := arr[i]
		var j int
		for j =i-1;j>=0 && arr[j] > t; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = t
	}
}
