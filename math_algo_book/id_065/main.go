package main

import (
	"fmt"
	"math"
)

func main() {
	var H, W int
	fmt.Scan(&H, &W)

	answer := int(math.Ceil(float64(H) * float64(W) / 2))
	fmt.Println(answer)
}
