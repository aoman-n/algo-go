package bubble

// func Sort(list []int) []int {
// 	for i := 0; i < len(list)-1; i++ {
// 		for j := 0; j < len(list)-i-1; j++ {
// 			if list[j] > list[j+1] {
// 				list[j], list[j+1] = list[j+1], list[j]
// 			}
// 		}
// 	}

// 	return list
// }

// func Sort(list []int) []int {
// 	for i := 0; i < len(list); i++ {
// 		for j := 0; j < len(list)-1-i; j++ {
// 			if list[j] > list[j+1] {
// 				list[j], list[j+1] = list[j+1], list[j]
// 			}
// 		}
// 	}

// 	return list
// }

func Sort(list []int) []int {
	for i := range make([]struct{}, len(list)) {
		for j := range make([]struct{}, len(list)-1-i) {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}

	return list
}
