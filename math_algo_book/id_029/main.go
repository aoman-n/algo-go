package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// n=3,[0,0,0,0]
	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}

	answer := dp[n]
	fmt.Println(answer)
}
