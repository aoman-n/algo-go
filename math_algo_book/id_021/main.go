package main

import "fmt"

func factorial(n int) int {
	sum := 1
	for i := 2; i <= n; i++ {
		sum *= i
	}
	return sum
}

// nCr
func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var n, r int
	fmt.Scan(&n, &r)

	fmt.Println(combination(n, r))
}
