package main

import (
	"fmt"
	"math"
)

// <問題文>
// N 個の品物があります。 品物には 1,2,…,N と番号が振られています。 各 i (1≤i≤N) について、品物 i の重さは wi​で、価値は viです。
// 太郎君は、N 個の品物のうちいくつかを選び、ナップサックに入れて持ち帰ることにしました。 ナップサックの容量は W であり、持ち帰る品物の重さの総和は W 以下でなければなりません。
// 太郎君が持ち帰る品物の価値の総和の最大値を求めてください。

// <入力>
// N W
// W1 V1
// W2 V2
// W3 V3
// ...
// Wi Vi
func main() {
	var N, W int
	fmt.Scan(&N, &W)

	w := make([]int, N+1)
	v := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Scan(&w[i], &v[i])
	}

	// dp[i][w]
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= N; i++ {
		for j := 0; j <= W; j++ {
			if w[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i]]+v[i])
			}
		}
	}

	answer := dp[N][W]
	fmt.Println(answer)
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
