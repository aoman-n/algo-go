package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var answer int
	switch n % 4 {
	case 0:
		answer = 6
	case 1:
		answer = 2
	case 2:
		answer = 4
	case 3:
		answer = 8
	}

	fmt.Println(answer)
}
