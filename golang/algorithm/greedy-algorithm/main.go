package main

import "fmt"

func main() {
	currencies := []int{100, 50, 20, 10, 1}
	smallChange := GiveChange(currencies, 272)
	fmt.Println(smallChange)
}

func GiveChange(currencies []int, price int) []int {
	var payments []int
	for price > 0 {
		currency := getHighestCurrency(currencies, price)
		price -= currency
		payments = append(payments, currency)
	}

	return payments
}

func getHighestCurrency(currencies []int, price int) int {
	for _, v := range currencies {
		if price >= v {
			return v
		}
	}
	return 0
}
