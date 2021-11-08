package main

import "fmt"

/*
母牛生小牛问题

有一只母牛，每年初生一只小母牛，每头小母牛从第三个年头起，每年生一只小母牛。
求指定年数有多少头牛？
*/
func main() {
	year := 20
	count := GetCow(year)
	fmt.Printf("%d年有%d只牛", year, count)
}

func GetCow(year int) int {
	// 3个元素分别代表第1年、第2年、第3年数量
	arr := []int{0, 0, 0}
	// 最先有一只小母牛
	arr[0] = 1
	for i := 1; i <= year; i++ {
		// 第三年的生下小牛数
		young := arr[2]
		// 第二年的母牛长到第三年
		arr[2] += arr[1]
		// 第一年的母年长到第二年
		arr[1] = arr[0]
		// 将新出生的放到第1年
		arr[0] = young
	}

	return arr[0] + arr[1] + arr[2]
}
