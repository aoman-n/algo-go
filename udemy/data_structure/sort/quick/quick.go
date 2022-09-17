package quick

// partition
// pivotを基準に、左辺->pivotより小さい数、右辺->pivotより大きい数に並べ替え、
// pivot挿入位置のindexを返す関数
func partition(numbers []int, row, high int) int {
	i := row - 1
	pivot := numbers[high]

	for j := row; j < high; j++ {
		if numbers[j] < pivot {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}

	numbers[i+1], numbers[high] = numbers[high], numbers[i+1]
	return i + 1
}

// 再帰的によび出すためのヘルパー
func sortHelper(numbers []int, row, high int) {
	if row < high {
		partitionIdx := partition(numbers, row, high)
		// 右辺
		sortHelper(numbers, row, partitionIdx-1)
		// 左辺
		sortHelper(numbers, partitionIdx+1, high)
	}
}

func Sort(input []int) []int {
	numbers := make([]int, len(input))
	copy(numbers, input)
	sortHelper(numbers, 0, len(numbers)-1)
	return numbers
}
