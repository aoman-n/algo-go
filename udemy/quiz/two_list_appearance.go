package quiz

func cacheCount(list []int) map[int]int {
	ret := make(map[int]int)
	for _, n := range list {
		ret[n] += 1
	}
	return ret
}

func MinCountRemove(x, y []int) ([]int, []int) {
	cacheX := cacheCount(x)
	cacheY := cacheCount(y)

	retX := make([]int, 0)
	retY := make([]int, 0)

	for _, n := range x {
		if cacheX[n] >= cacheY[n] {
			retX = append(retX, n)
		}
	}

	for _, n := range y {
		if cacheY[n] >= cacheX[n] {
			retY = append(retY, n)
		}
	}

	return retX, retY
}
