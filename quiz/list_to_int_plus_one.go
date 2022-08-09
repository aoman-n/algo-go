package quiz

import "math"

func ListToIntPlusOne(numbers []int) int {
	sum := 0
	i := len(numbers) - 1
	for i >= 0 {
		targetNum := numbers[i]
		multi := math.Pow(10, float64(len(numbers)-1-i))
		sum += targetNum * int(multi)
		i--
	}

	return sum + 1
}
