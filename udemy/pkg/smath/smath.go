package smath

func MaxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func MinInt(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func SumInts(ints []int) int {
	total := 0
	for _, v := range ints {
		total += v
	}
	return total
}

func MaxInNumbers(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > result {
			result = numbers[i]
		}
	}
	return result
}
