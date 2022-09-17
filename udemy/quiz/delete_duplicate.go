package quiz

import (
	"github.com/laster18/algo-go/udemy/pkg/slice"
)

// [1, 3, 3, 5, 5, 7, 7, 7, 10, 12, 12, 15]

// 別でsliceを作って、最後に書き換える
// for使用
func DeleteDuplicateV1(numbers *[]int) {
	tmp := []int{}
	for i := 0; i < len(*numbers); i++ {
		if !slice.Exists(tmp, (*numbers)[i]) {
			tmp = append(tmp, (*numbers)[i])
		}
	}
	*numbers = tmp
}

// 別でsliceを作って、最後に書き換える
// while使用
func DeleteDuplicateV2(numbers *[]int) {
	tmp := []int{(*numbers)[0]}

	i, lenNum := 0, len(*numbers)-1
	for i < lenNum {
		if (*numbers)[i] != (*numbers)[i+1] {
			tmp = append(tmp, (*numbers)[i+1])
		}
		i++
	}

	*numbers = tmp
}

// while使用して引数を直接書き換える
// 先頭からループ
func DeleteDuplicateV3(numbers *[]int) {
	i := 0
	for i < len(*numbers)-1 {
		if (*numbers)[i] == (*numbers)[i+1] {
			slice.Remove(numbers, i)
			i--
		}
		i++
	}
}

// while使用して引数を直接書き換える
// 最後からループ
func DeleteDuplicateV4(numbers *[]int) {
	for i := len(*numbers) - 1; i > 0; i-- {
		if (*numbers)[i] == (*numbers)[i-1] {
			slice.Remove(numbers, i)
		}
	}
}
