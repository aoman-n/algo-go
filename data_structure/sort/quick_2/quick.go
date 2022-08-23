package quick

// func partition(list []int, low, high int) int {
// 	pivot := list[high]
// 	i := low - 1
// 	j := low

// 	for j < high {
// 		if list[j] < pivot {
// 			i++
// 			list[i], list[j] = list[j], list[i]
// 		}
// 		j++
// 	}

// 	list[i+1], list[high] = list[high], list[i+1]
// 	return i + 1
// }

// partitionを利用した方法
// func Sort(list []int) []int {
// 	newList := make([]int, len(list))
// 	copy(newList, list)

// 	var helper func(low, high int)

// 	helper = func(low, high int) {
// 		if low < high {
// 			partitionIndex := partition(newList, low, high)
// 			// パーティションよりleft側
// 			helper(low, partitionIndex-1)
// 			// パーティションよりright側
// 			helper(partitionIndex+1, high)
// 		}
// 	}

// 	helper(0, len(newList)-1)

// 	return newList
// }

// leftとrightを新たにメモリアロケートする方法。メモリ効率は悪いが簡単
// func Sort(list []int) []int {
// 	newList := make([]int, len(list))
// 	copy(newList, list)

// 	if len(newList) < 2 {
// 		return newList
// 	}

// 	pivot := newList[0]
// 	left := []int{}
// 	right := []int{}

// 	for i := 1; i < len(newList); i++ {
// 		if newList[i] < pivot {
// 			left = append(left, newList[i])
// 		} else {
// 			right = append(right, newList[i])
// 		}
// 	}

// 	left = Sort(left)
// 	right = Sort(right)

// 	result := make([]int, len(left)+len(right)+1)
// 	copy(result[:len(left)], left)
// 	copy(result[len(left)+1:], right)
// 	result[len(left)] = pivot
// 	return result
// }

// return int: パーティションのインデックスを返す
func partition(list []int, low, high int) int {
	pivot := list[high]
	i := low - 1
	j := low

	for j < high {
		if list[j] < pivot {
			i++
			list[i], list[j] = list[j], list[i]
		}
		j++
	}

	list[i+1], list[high] = list[high], list[i+1]
	return i + 1
}

func Sort(list []int) []int {
	newList := make([]int, len(list))
	copy(newList, list)

	var sortHelper func(low, high int)

	sortHelper = func(low, high int) {
		if low < high {
			partitionIndex := partition(newList, low, high)
			sortHelper(low, partitionIndex-1)
			sortHelper(partitionIndex+1, high)
		}

	}

	sortHelper(0, len(list)-1)

	return newList
}
