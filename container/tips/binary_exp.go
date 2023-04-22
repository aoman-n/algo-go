// You can edit this code!
// Click here and start typing.
package main

import (
	"math"
)

func getMinBinaryExp(x int) int {
	// 10^9 < 2^n 時の最小のnを調べる
	l := math.Pow(10, float64(x))

	n := 0
	for {
		if l <= math.Pow(2, float64(n)) {
			break
		}
		n++
	}

	return n
}
