package main

import "fmt"

// 1 <= x, y <= n
// x+y <= s
func main() {
	var n, s int
	fmt.Scan(&n, &s)

	var count int
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			sum := i + j
			if sum <= s {
				count++
			}
		}
	}

	fmt.Println(count)
}
