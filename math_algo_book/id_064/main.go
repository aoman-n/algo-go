package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	A := make([]int, N)
	for i := range A {
		fmt.Scan(&A[i])
	}

	sum := 0
	for _, v := range A {
		sum += v
	}

	if sum%2 != K%2 {
		fmt.Println("No")
	} else if K < sum {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
