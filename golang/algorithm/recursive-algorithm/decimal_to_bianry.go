package main

import (
	"fmt"
	"math"
)

// 递归：
// 10进制转2进制
func decimalToBinary(num int64) int64 {
	// 结果
	var result int64 = 0
	// 保存位置
	var pos float64 = 0

	for num > 0 {
		a := num % 2
		num = num / 2
		result = int64(math.Pow(10, pos))*a + result
		pos += 1
	}

	return result
}

func main() {
	var num int64 = 123
	fmt.Printf("%d\n", decimalToBinary(num))
}
