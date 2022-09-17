package main

import "fmt"

const (
	First  = "First"
	Second = "Second"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n%4 == 0 {
		fmt.Println(Second)
	} else {
		fmt.Println(First)
	}
}
