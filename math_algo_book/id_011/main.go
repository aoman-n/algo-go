package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 一旦O(n**2)
func main() {
	var n int
	fmt.Scan(&n)

	primes := make([]string, 0)
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, strconv.Itoa(i))
		}
	}

	fmt.Println(strings.Join(primes, " "))
}

func isPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
