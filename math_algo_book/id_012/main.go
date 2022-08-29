package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	answer := "Yes"
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			answer = "No"
			break
		}
	}

	fmt.Println(answer)
}
