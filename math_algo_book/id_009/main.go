package main

import (
	"fmt"
	"math"
)

// Bad: ビット全探索
func main() {
	var n, s int
	fmt.Scan(&n, &s)

	list := make([]int, n)
	for i := range list {
		fmt.Scan(&list[i])
	}

	for i := 0; i < int(math.Pow(2, float64(n))); i++ {
		sum := 0

		for j := 1; j <= n; j++ {
			validDigit := (i & int(math.Pow(2, float64(i-j)))) == 0
			if validDigit {
				sum += list[j-1]
			}
		}

		if sum == s {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
