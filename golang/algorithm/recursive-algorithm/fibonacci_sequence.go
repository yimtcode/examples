package main

import "fmt"

// 斐波那契数列
func fibonacciSequence(n int) int {
	if n < 0 {
		panic("n value illegal")
	}

	if n < 2 {
		return n
	}

	return fibonacciSequence(n-1) + fibonacciSequence(n-2)
}

func main() {
	num := fibonacciSequence(8)
	fmt.Println(num)
}
