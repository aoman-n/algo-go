package hashtable

func GetPair(numbers []int, target int) (int, int, bool) {
	m := make(map[int]bool)

	for _, n := range numbers {
		if ok := m[target-n]; ok {
			return target - n, n, true
		}

		m[n] = true
	}

	return 0, 0, false
}

func GetPairHalfSum(numbers []int) (int, int, bool) {
	sum := 0
	for _, d := range numbers {
		sum += d
	}

	if sum%2 != 0 {
		return 0, 0, false
	}

	m := make(map[int]bool)
	half := sum / 2

	for _, n := range numbers {
		if ok := m[half-n]; ok {
			return half - n, n, true
		}

		m[n] = true
	}

	return 0, 0, false
}
