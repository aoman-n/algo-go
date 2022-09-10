package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}

	dp := make([]int, n+1)

	dp[2] = abs(a[2] - a[1])

	for i := 3; i <= n; i++ {
		dp[i] = min(
			dp[i-2]+abs(a[i]-a[i-2]),
			dp[i-1]+abs(a[i]-a[i-1]),
		)
	}

	answer := dp[n]
	fmt.Println(answer)
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
