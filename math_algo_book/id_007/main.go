package main

import "fmt"

func main() {
	var n int
	var x, y int

	fmt.Scan(&n)
	fmt.Scan(&x)
	fmt.Scan(&y)

	var count int
	for i := 1; i <= n; i++ {
		if (i%x == 0) || (i%y == 0) {
			count++
		}
	}
	fmt.Println(count)
}
