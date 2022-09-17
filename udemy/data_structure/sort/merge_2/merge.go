package merge

import "math"

func Sort(list []int) []int {
	newList := make([]int, len(list))
	copy(newList, list)
	return sortHelper(newList)
}

func sortHelper(list []int) []int {
	if len(list) == 1 {
		return []int{list[0]}
	}

	mid := len(list) / 2

	left := sortHelper(list[:mid])
	right := sortHelper(list[mid:])

	totalLen := len(left) + len(right)
	left = append(left, math.MaxInt)
	right = append(right, math.MaxInt)

	result := make([]int, 0, totalLen)
	leftIndex := 0
	rightIndex := 0
	for len(result) < totalLen {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	return result
}
