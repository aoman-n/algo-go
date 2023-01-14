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

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	N, S := ni2(sc)
	A := make([]int, N+1)
	for i := 1; i < N+1; i++ {
		A[i] = ni(sc)
	}

	dp := make([][]bool, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]bool, S+1)
	}

	dp[0][0] = true

	// i=縦軸のindex
	for i := 0; i < N; i++ {
		targetI := i + 1
		// j=横軸のindex
		for j := 0; j < S+1; j++ {
			if dp[i][j] {
				dp[targetI][j] = true
				sum := j + A[targetI]
				if sum < S+1 {
					dp[targetI][j+A[targetI]] = true
				}
			}
		}
	}

	if dp[N][S] {
		fmt.Fprint(writer, "Yes")
	} else {
		fmt.Fprint(writer, "No")
	}
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

func ni3(sc *bufio.Scanner) (int, int, int) {
	return ni(sc), ni(sc), ni(sc)
}

func ni4(sc *bufio.Scanner) (int, int, int, int) {
	return ni(sc), ni(sc), ni(sc), ni(sc)
}

func nis(sc *bufio.Scanner, n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni(sc)
	}
	return a
}

// n: 行数
// m: 要素数
func nir(sc *bufio.Scanner, n int, m int) [][]int {
	a := make([][]int, n)
	for i := range a {
		a[i] = nis(sc, m)
	}
	return a
}

// ==================================================
// util
// ==================================================

func abs(n int) int {
	return int(math.Abs(float64(n)))
}
