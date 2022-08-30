package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	numbers := make([]int, n)
	for i := range numbers {
		fmt.Scan(&numbers[i])
	}

	answer := leastCommonMultiple(numbers[0], numbers[1])
	for i := 2; i < len(numbers); i++ {
		answer = leastCommonMultiple(answer, numbers[i])
	}

	fmt.Println(answer)
}

func leastCommonMultiple(x, y int) int {
	return (x / greatestCommonDivisor(x, y)) * y
}

func greatestCommonDivisor(x, y int) int {
	if y == 0 {
		return x
	}

	return greatestCommonDivisor(y, x%y)
}
