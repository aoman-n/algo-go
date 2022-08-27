package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	answer := n
	for i := n - 1; i > 0; i-- {
		answer *= i
	}

	fmt.Println(answer)
}
