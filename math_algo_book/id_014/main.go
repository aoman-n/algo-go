package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 解1
// func divide(n int) (divisor int, ok bool) {
// 	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
// 		if n%i == 0 {
// 			return i, true
// 		}
// 	}

// 	return 0, false
// }

// func main() {
// 	var n int
// 	fmt.Scan(&n)

// 	divisors := make([]string, 0)
// 	for {
// 		divisor, ok := divide(n)
// 		if ok {
// 			divisors = append(divisors, strconv.Itoa(divisor))
// 			n = n / divisor
// 		} else {
// 			divisors = append(divisors, strconv.Itoa(n))
// 			break
// 		}
// 	}

// 	fmt.Println(strings.Join(divisors, " "))
// }

// 解2
func main() {
	var n int
	fmt.Scan(&n)

	answer := make([]string, 0)
	limit := int(math.Sqrt(float64(n)))

	for i := 2; i <= limit; i++ {
		for {
			if n%i == 0 {
				answer = append(answer, strconv.Itoa(i))
				n = n / i
			} else {
				break
			}
		}
	}

	fmt.Println(strings.Join(answer, " "))
}
