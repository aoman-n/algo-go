package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]float64, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	b := make([]float64, n)
	for i := range a {
		fmt.Scan(&b[i])
	}

	// 合計勉強時間の期待値を求める

	// 1. 日の最初にサイコロを振る
	// 2. サイコロを振って 1,2 が出た場合: Ai 時間勉強する
	// 3. サイコロを振って 3,4,5,6 が出た場合: Bi 時間勉強する

	// (2/6 * Ai) + (4/6 * Bi)

	var answer float64
	for i := 0; i < n; i++ {
		answer += (float64(2) / float64(6) * a[i]) + (float64(4) / float64(6) * b[i])
	}

	fmt.Println(answer)
}
