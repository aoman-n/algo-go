package insertion

// func Sort(list []int) []int {
// 	result := make([]int, len(list))
// 	copy(result, list)

// 	for i := 1; i < len(result); i++ {
// 		for j := i; j > 0; j-- {
// 			if result[j] < result[j-1] {
// 				result[j], result[j-1] = result[j-1], result[j]
// 			} else {
// 				break
// 			}
// 		}

// 	}

// 	return result
// }

// func Sort(list []int) []int {
// 	newList := make([]int, len(list))
// 	copy(newList, list)

// 	for i := 1; i < len(newList); i++ {
// 		tmp := newList[i]
// 		j := i - 1
// 		for j >= 0 && newList[j] > tmp {
// 			newList[j+1] = newList[j]
// 			j--
// 		}

// 		newList[j+1] = tmp
// 	}

// 	return newList
// }

func Sort(list []int) []int {
	newList := make([]int, len(list))
	copy(newList, list)

	for i := 1; i < len(newList); i++ {
		tmp := newList[i]
		j := i - 1
		for j >= 0 && newList[j] > tmp {
			newList[j+1] = newList[j]
			j--
		}
		newList[j+1] = tmp
	}

	return newList
}
