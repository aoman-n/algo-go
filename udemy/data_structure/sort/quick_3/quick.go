package quick

// func partition(list []int, start, end int) int {
// 	pivot := list[end]
// 	i := start - 1
// 	j := start
// 	for j < end {
// 		if list[j] < pivot {
// 			i++
// 			list[i], list[j] = list[j], list[i]
// 		}
// 		j++
// 	}
// 	list[end], list[i+1] = list[i+1], list[end]
// 	return i + 1
// }

// func Sort(list []int) []int {
// 	newList := make([]int, len(list))
// 	copy(newList, list)

// 	var helper func(left, right int)

// 	helper = func(start, end int) {
// 		if start < end {
// 			partitionIndex := partition(newList, start, end)
// 			helper(start, partitionIndex-1)
// 			helper(partitionIndex+1, end)
// 		}
// 	}

// 	helper(0, len(newList)-1)
// 	return newList
// }

func Sort(list []int) []int {
	newList := make([]int, len(list))
	copy(newList, list)
	SortHelper(newList, 0, len(newList)-1)
	return newList
}

func SortHelper(list []int, start, end int) {
	mid := (start + end) / 2
	pivot := list[mid]
	i := start
	j := end

	for i <= j {
		for list[i] < pivot {
			i++
		}
		for list[j] > pivot {
			j--
		}
		if i >= j {
			break
		}

		list[i], list[j] = list[j], list[i]
	}

	if start < i-1 {
		SortHelper(list, start, i-1)
	}
	if end > j+1 {
		SortHelper(list, j+1, end)
	}
}
