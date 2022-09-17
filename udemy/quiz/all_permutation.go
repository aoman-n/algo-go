package quiz

func insertForPerm(list []int, i int, val int) []int {
	result := make([]int, len(list))
	copy(result, list)
	result = append(result, 0)
	copy(result[i+1:], result[i:])
	result[i] = val
	return result
}

func AllPermutation(elements []int) [][]int {
	result := [][]int{}

	if len(elements) <= 1 {
		return [][]int{elements}
	}

	for _, perm := range AllPermutation(elements[1:]) {
		for i := 0; i < len(elements); i++ {
			s := insertForPerm(perm, i, elements[0])
			result = append(result, s)
		}
	}

	return result
}
