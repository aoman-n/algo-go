package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	numbers := make([]int, n)
	for i := range numbers {
		fmt.Scan(&numbers[i])
	}

	for i := 0; i < len(numbers)-1; i++ {
		x := numbers[i]
		y := numbers[i+1]
		divisor := greatestCommonDivisor(x, y)
		numbers[i+1] = divisor
	}

	answer := numbers[len(numbers)-1]
	fmt.Println(answer)
}

func greatestCommonDivisor(x, y int) int {
	if y == 0 {
		return x
	}

	return greatestCommonDivisor(y, x%y)
}
