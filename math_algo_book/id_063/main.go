package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var answer string
	if n%2 != 0 {
		answer = "No"
	} else {
		answer = "Yes"
	}

	fmt.Println(answer)
}
