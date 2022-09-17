package generator

import (
	"math/rand"
	"time"
)

func Ints(len int) []int {
	rand.Seed(time.Now().UnixNano())
	ints := make([]int, len)
	for i := range ints {
		ints[i] = rand.Intn(1000)
	}
	return ints
}
