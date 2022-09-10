package main

import (
	"fmt"
	"math"
)

// FIXME: 下記コードは間に合わないので改善の必要あり
func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}

	// i日目に勉強する場合
	dp1 := make([]int, n+1)
	// i日目に勉強しない場合
	dp2 := make([]int, n+1)

	dp1[0] = 0
	dp2[0] = 0

	for i := 1; i <= n; i++ {
		dp1[i] = dp2[i-1] + a[i]
		dp2[i] = max(dp1[i-1], dp2[i-1])
	}

	answer := max(dp1[n], dp2[n])
	fmt.Println(answer)
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
