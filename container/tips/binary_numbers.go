package tips

import (
	"fmt"
	"math"
)

func binaryNumbers() {
	n := 1000000000
	// n := 126

	ints := []int{}

	for {
		if n == 0 {
			break
		}

		ints = append(ints, n%2)

		n = n / 2
	}

	// reverse
	for i := len(ints)/2 - 1; i >= 0; i-- {
		opp := len(ints) - i - 1
		ints[i], ints[opp] = ints[opp], ints[i]
	}

	// 2進数の各桁のビットが表示される
	fmt.Printf("ints(%v)\n", ints)
	fmt.Printf("ints.len(%v)\n", len(ints))

	nums := make([]int, 0)

	for i := 0; i < 30; i++ {
		nums = append(nums, int(math.Pow(2, float64(i))))
	}

	fmt.Printf("nums(%v)\n", nums)

}
