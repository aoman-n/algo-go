package main

import "fmt"

// // Bad: ビット全探索
// func main() {
// 	var n, s int
// 	fmt.Scan(&n, &s)

// 	list := make([]int, n)
// 	for i := range list {
// 		fmt.Scan(&list[i])
// 	}

// 	for i := 0; i < int(math.Pow(2, float64(n))); i++ {
// 		sum := 0

// 		for j := 1; j <= n; j++ {
// 			validDigit := (i & int(math.Pow(2, float64(i-j)))) == 0
// 			if validDigit {
// 				sum += list[j-1]
// 			}
// 		}

// 		if sum == s {
// 			fmt.Println("Yes")
// 			return
// 		}
// 	}

// 	fmt.Println("No")
// }

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}

	dp := make([][]bool, n+1)
	for i := range make([]struct{}, n+1) {
		dp[i] = make([]bool, s+1)
	}

	dp[0][0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j <= s; j++ {
			if a[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-a[i]]
			}
		}
	}

	if dp[n][s] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
