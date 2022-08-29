package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	divisors := make([]string, 0)

	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			divisors = append(divisors, strconv.Itoa(i), strconv.Itoa(n/i))
		}
	}

	fmt.Println(strings.Join(divisors, "\n"))
}
