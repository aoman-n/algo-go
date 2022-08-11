package binary

type IndexInt = int

func Search(s []int, v int) IndexInt {
	if len(s) <= 0 {
		return -1
	}

	return searchHelper(s, v, 0, len(s)-1)
}

func searchHelper(s []int, v int, left, right int) IndexInt {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	if s[mid] == v {
		return mid
	}

	if v < s[mid] {
		return searchHelper(s, v, left, mid-1)
	} else {
		return searchHelper(s, v, mid+1, right)
	}
}
