package selection

func Sort(ints []int) []int {
	numbers := make([]int, len(ints))
	copy(numbers, ints)

	for i := 0; i < len(ints); i++ {
		tmpIdx := i
		for j := i + 1; j < len(ints); j++ {
			if numbers[tmpIdx] > numbers[j] {
				tmpIdx = j
			}
		}
		numbers[i], numbers[tmpIdx] = numbers[tmpIdx], numbers[i]
	}

	return numbers
}
