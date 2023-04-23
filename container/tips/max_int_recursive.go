package tips

import "math"

// ------------------------------------
// ------------------------------------

// 半開区間を利用した再帰処理
func GetMax(s []int) int {
	return getMaxHelper(s, 0, len(s))
}

func getMaxHelper(s []int, l, r int) int {
	if r-l == 1 {
		return s[l]
	}

	mid := (r + l) / 2
	lMax := getMaxHelper(s, l, mid)
	rMax := getMaxHelper(s, mid, r)

	return int(math.Max(float64(lMax), float64(rMax)))
}

// ------------------------------------
// ------------------------------------

// func GetMax(s []int) int {
// 	return getMaxHelper(s, 0, len(s)-1)
// }

// func getMaxHelper(s []int, l, r int) int {
// 	if l == r {
// 		return s[l]
// 	}

// 	mid := (l + r) / 2
// 	lMax := getMaxHelper(s, l, mid)
// 	rMax := getMaxHelper(s, mid+1, r)

// 	return int(math.Max(float64(lMax), float64(rMax)))
// }
