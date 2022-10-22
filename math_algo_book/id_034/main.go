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

const MOD = 1000000007

func init() {
	reader = os.Stdin
	writer = os.Stdout
}

// O(N**2)の計算量のパターン
func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	N := ni(sc)

	X := make([]int, N+1)
	Y := make([]int, N+1)
	for i := 1; i <= N; i++ {
		X[i], Y[i] = ni2(sc)
	}

	answer := float64(10000000000000)
	for i := 1; i < N; i++ {
		for j := i + 1; j <= N; j++ {
			iX, iY := X[i], Y[i]
			jX, jY := X[j], Y[j]
			// (iX, iY) -> (jX, jY)
			// v(jX-iX, jY-iY)
			vX, vY := jX-iX, jY-iY
			v := math.Abs(math.Sqrt((float64(vX) * float64(vX)) + (float64(vY) * float64(vY))))
			// fmt.Printf("(%v, %v) -> (%v, %v): ", iX, iY, jX, jY)
			// fmt.Printf("v: %v\n", v)
			if v < answer {
				answer = v
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
