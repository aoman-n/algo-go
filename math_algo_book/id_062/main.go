package main

import (
	"fmt"
)

// k回のテレポート
// n個の町
// 町i->町Aiへのテレポート
// 町1からスタートしてk回テレポート
func main() {
	var N, K int
	fmt.Scan(&N, &K)

	A := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Scan(&A[i])
	}

	currentTown := 1
	transitionCount := 0
	visitedDowns := []int{currentTown}
	visitedDownsMap := make(map[int]bool)
	visitedDownsMap[currentTown] = true

	var answer int
	for transitionCount <= K {
		currentTown = A[currentTown]
		transitionCount++
		if transitionCount == K {
			answer = currentTown
			break
		}

		if visitedDownsMap[currentTown] {
			foundIndex := findIndex(visitedDowns, currentTown)
			loopCount := len(visitedDowns) - foundIndex
			loopTowns := visitedDowns[foundIndex:]
			answer = loopTowns[(K-foundIndex)%loopCount]
			break
		} else {
			visitedDowns = append(visitedDowns, currentTown)
			visitedDownsMap[currentTown] = true
		}
	}

	fmt.Println(answer)
}

func findIndex(s []int, target int) int {
	for i, n := range s {
		if n == target {
			return i
		}
	}
	return -1
}
