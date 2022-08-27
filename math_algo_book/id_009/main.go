package main

import "fmt"

// Bad: 全探索(未完成)
func main() {
	var n, s int
	fmt.Scan(&n, &s)

	list := make([]int, n)
	for i := range list {
		fmt.Scan(&list[i])
	}

	answer := "No"
loop:
	for i, x := range list {
		for j := i + 1; j < len(list); j++ {
			y := list[j]
			fmt.Printf("x: %v, y: %v\n", x, y)
			sum := x + y
			if sum == s {
				answer = "Yes"
				break loop
			}
		}
	}

	fmt.Println(answer)
}
