package bubble

// func genRandomInts(len int) []int {
// 	rand.Seed(time.Now().UnixNano())
// 	ret := make([]int, len)
// 	for i := range ret {
// 		ret[i] = rand.Intn(1000)
// 	}
// 	return ret
// }

func Sort(ints []int) []int {
	for i := 0; i < len(ints)-1; i++ {
		for j := 0; j < len(ints)-1-i; j++ {
			if ints[j] > ints[j+1] {
				ints[j], ints[j+1] = ints[j+1], ints[j]
			}
		}
	}
	return ints
}
