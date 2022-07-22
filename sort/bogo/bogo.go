package bogo

import (
	"math/rand"
	"time"
)

func randomInts(len int) []int {
	rand.Seed(time.Now().UnixNano())
	ret := make([]int, len)
	for i := range ret {
		ret[i] = rand.Intn(1000)
	}
	return ret
}

func shuffle(list []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
}

func inOrder(numbers []int) bool {
	for i := range make([]interface{}, len(numbers)-1) {
		if numbers[i] > numbers[i+1] {
			return false
		}
	}
	return true
}

func Sort(numbers []int) {
	for !inOrder(numbers) {
		shuffle(numbers)
	}
}
