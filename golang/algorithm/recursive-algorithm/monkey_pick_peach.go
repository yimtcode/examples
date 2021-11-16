package main

import "fmt"

/*
一只小猴子一天摘了许多桃子，第一天吃了一半，然后忍不住又吃了一个；
第二天又吃了一半，再加上一个；后面每天都是这样吃。到第10天的时候，
小猴子发现只有一个桃子了。问小猴子第一天共摘了多少个桃子。
*/
func main() {
	day := 10
	y := 1
	count := getPeach(day, y)
	fmt.Println(count)
}

func getPeach(day, y int) int {
	x := y
	for i := day; i > 1; i-- {
		x += 1
		x *= 2
	}
	return x
}
