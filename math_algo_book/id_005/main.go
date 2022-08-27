package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	var sum int
	for _, num := range a {
		sum += num
	}

	answer := sum % 100

	fmt.Println(answer)
}
