package merge

func sortHelper(arr []int, start, end int) []int {
	if start == end {
		return []int{arr[start]}
	}

	middle := (end + start) / 2
	left := sortHelper(arr, start, middle)
	right := sortHelper(arr, middle+1, end)

	merged := make([]int, len(left)+len(right))

	li, ri, mi := 0, 0, 0
	for li < len(left) && ri < len(right) {
		if left[li] < right[ri] {
			merged[mi] = left[li]
			li++
		} else {
			merged[mi] = right[ri]
			ri++
		}
		mi++
	}

	for li < len(left) {
		merged[mi] = left[li]
		li++
		mi++
	}

	for ri < len(right) {
		merged[mi] = right[ri]
		ri++
		mi++
	}

	return merged
}

func Sort(arr []int) []int {
	if len(arr) <= 0 {
		return arr
	}

	copied := make([]int, len(arr))
	copy(copied, arr)

	return sortHelper(copied, 0, len(copied)-1)
}
