package selection

// func Sort(list []int) []int {
// 	for i := 0; i < len(list)-1; i++ {
// 		var minIndex = i
// 		for j := i + 1; j < len(list); j++ {
// 			if list[j] < list[minIndex] {
// 				minIndex = j
// 			}
// 		}
// 		list[i], list[minIndex] = list[minIndex], list[i]
// 	}

// 	return list
// }

func Sort(list []int) []int {
	sortedList := make([]int, len(list))
	copy(sortedList, list)

	for i := range make([]struct{}, len(sortedList)) {
		minIndex := i
		for j := i + 1; j < len(sortedList); j++ {
			if sortedList[j] < sortedList[minIndex] {
				minIndex = j
			}
		}
		sortedList[i], sortedList[minIndex] = sortedList[minIndex], sortedList[i]
	}

	return sortedList
}
