package main

import "fmt"

// 2枚選ぶ
// 総和=100000 は何通り

// 2≤N≤200000
// 1≤Ai≤99999
//      50000
// 入力はすべて整数

func main() {
	var n int
	fmt.Scan(&n)

	counts := map[int]int{}
	for range make([]struct{}, n) {
		var a int
		fmt.Scan(&a)
		counts[a]++
	}

	var answer int
	for i := 1; i <= 49999; i++ {
		answer += counts[i] * counts[100000-i]
	}
	answer += counts[50000] * (counts[50000] - 1) / 2

	fmt.Println(answer)
}
