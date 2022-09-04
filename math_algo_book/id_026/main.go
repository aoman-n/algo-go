package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)

	var answer float64
	for i := 1; float64(i) <= n; i++ {
		num := math.Pow(float64(1/n), float64(i))
		fmt.Printf("num: %v, i: %v\n", num, i)
		answer += math.Pow(float64(1/n), float64(i))
	}

	fmt.Println(answer)
}
