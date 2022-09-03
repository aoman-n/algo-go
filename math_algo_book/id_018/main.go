package main

import "fmt"

// pricesは「100,200,300,400」のいずれか
// 2つ選んで、合計500円になるのは何通りか？

// 1. 100 + 400 = 500
// 2. 200 + 300 = 500
func main() {
	var n int
	fmt.Scan(&n)

	counts := map[int]int{
		100: 0,
		200: 0,
		300: 0,
		400: 0,
	}

	for range make([]struct{}, n) {
		var a int
		fmt.Scan(&a)
		counts[a]++
	}

	answer := (counts[100] * counts[400]) + (counts[300] * counts[200])
	fmt.Println(answer)
}
