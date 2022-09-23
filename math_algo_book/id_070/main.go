package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

var reader io.Reader
var writer io.Writer

func init() {
	reader = os.Stdin
	writer = os.Stdout
}

func checkNumPoints(left, right, top, bottom int, pointX, pointY []int, minCount int) bool {
	counter := 0
	for i := 1; i < len(pointX); i++ {
		inX := left <= pointX[i] && pointX[i] <= right
		inY := bottom <= pointY[i] && pointY[i] <= top

		if inX && inY {
			counter++

			if counter >= minCount {
				return true
			}
		}
	}

	return false
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	N, K := ni2(sc)

	X := make([]int, N+1)
	Y := make([]int, N+1)

	for i := 1; i <= N; i++ {
		X[i], Y[i] = ni2(sc)
	}

	initialized := false
	answer := 0
	// (i, j, k, l) = (x左, x右, y上, y下)
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			for k := 1; k <= N; k++ {
				for l := 1; l <= N; l++ {
					left := X[i]   // x左
					right := X[j]  // x右
					top := Y[k]    // y上
					bottom := Y[l] // y下

					if checkNumPoints(left, right, top, bottom, X, Y, K) {
						area := (abs(right - left)) * (abs(top - bottom))
						if initialized {
							answer = min(answer, area)
						} else {
							answer = area
							initialized = true
						}
					}
				}
			}
		}
	}

	fmt.Fprint(writer, answer)
}

// ==================================================
// io
// ==================================================

func ni(sc *bufio.Scanner) int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func ni2(sc *bufio.Scanner) (int, int) {
	return ni(sc), ni(sc)
}
