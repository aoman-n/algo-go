package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	answer := helper(a, b)

	fmt.Println(answer)
}

func helper(a, b int) int {
	if b == 0 {
		return a
	}

	return helper(b, a%b)
}
