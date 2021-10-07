package main

import "fmt"

func main() {
	fmt.Println("示例1:")
	fmt.Println(substring("vista", "hish"))
	fmt.Println(substring("fish", "hish"))

	fmt.Println("示例2:")
	fmt.Println(subsequence("fish", "fosh"))
	fmt.Println(subsequence("fort", "fosh"))
}

func createMatrix(rows, cols int) [][]int {
	cell := make([][]int, rows)
	for i := range cell {
		cell[i] = make([]int, cols)
	}

	return cell
}

// 示例1
// 相同子串
func substring(a, b string) string {
	lcs := 0
	lastSubIndex := 0
	cell := createMatrix(len(a)+1, len(b)+1)

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				cell[i][j] = cell[i-1][j-1] + 1

				if cell[i][j] > lcs {
					lcs = cell[i][j]
					lastSubIndex = i
				}
			} else {
				cell[i][j] = 0
			}
		}
	}

	return a[lastSubIndex-lcs : lastSubIndex]
}

// 示例2
// 相同字符数
func subsequence(a, b string) int {
	cell := createMatrix(len(a)+1, len(b)+1)

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				cell[i][j] = cell[i-1][j-1] + 1
			} else {
				cell[i][j] = cell[i-1][j]

				if cell[i][j] < cell[i][j-1] {
					cell[i][j] = cell[i][j-1]
				}
			}
		}
	}

	return cell[len(a)][len(b)]
}
