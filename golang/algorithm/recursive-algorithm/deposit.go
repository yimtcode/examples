package main

import "fmt"

/*
问题:
		小明打算为自己的三年研究生生活准备一笔学费，一次性存入银行，
	保证在每年年底取出1000元，到第3年学习结束时刚好取完（月息0.31%）。
思路:
	3year = 1000 = 2year * (1 + 0.0031 * 12)
	2year = 1year * (1 + 0.0031 * 12)
	1year = deposit * (1 + 0.0031 * 12)
*/
func main() {
	year := 3
	deposit := 0.0
	for i := year; i >= 1; i-- {
		deposit = (deposit + 1000)/ (1 + 0.0031 * 12)
	}

	fmt.Printf("%0.2f\n", deposit)
}

