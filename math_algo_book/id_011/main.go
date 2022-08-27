package main

import (
	"fmt"
	"strconv"
	"strings"
)

// // 一旦O(n**2)
// func main() {
// 	var n int
// 	fmt.Scan(&n)

// 	primes := make([]string, 0)
// 	for i := 2; i <= n; i++ {
// 		if isPrime(i) {
// 			primes = append(primes, strconv.Itoa(i))
// 		}
// 	}

// 	fmt.Println(strings.Join(primes, " "))
// }

// func isPrime(x int) bool {
// 	for i := 2; i < x; i++ {
// 		if x%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// cacheを使った場合
func main() {
	var n int
	fmt.Scan(&n)

	isPrime := genIsPrimeFn()

	result := make([]string, 0)
	for i := 2; i <= n; i++ {
		if isPrime(i, n) {
			result = append(result, strconv.Itoa(i))
		}
	}

	fmt.Println(strings.Join(result, " "))
}

func genIsPrimeFn() func(int, int) bool {
	cache := map[int]bool{}

	return func(x, n int) bool {
		if cache[x] {
			return false
		}

		isPrime := true
		for i := 2; i < x; i++ {
			if x%i == 0 {
				isPrime = false
				break
			}
		}

		for i := x; i < n; i = i * 2 {
			cache[i] = true
		}

		return isPrime
	}
}

// 10
// 1, 2, 3, 4, 5, 6, 7, 8, 9

// target: 4
// 2, 3, 4

// -> 2, 3, 5, 7
