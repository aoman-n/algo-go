package insertion

func Sort(input []int) []int {
	ints := make([]int, len(input))
	copy(ints, input)

	for i := range ints {
		tmp := ints[i]
		j := i - 1

		for j >= 0 && ints[j] > tmp {
			ints[j+1] = ints[j]
			j--
		}
		ints[j+1] = tmp
	}

	return ints
}
