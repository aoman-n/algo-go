package quiz

type Pair struct {
	first  int
	second int
}

func FindSymmetricPairs(pairs []Pair) []Pair {
	m := make(map[int]int)
	ret := make([]Pair, 0)

	for _, data := range pairs {
		first := data.first
		second := data.second

		val, ok := m[second]
		if !ok {
			m[first] = second
			continue
		}

		if first == val {
			ret = append(ret, Pair{first, second})
		}
	}

	return ret
}
