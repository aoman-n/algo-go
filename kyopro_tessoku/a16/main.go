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
	N := ni(sc)

	A := make([]int, N)
	B := make([]int, N)

	for i := 1; i < N; i++ {
		A[i] = ni(sc)
	}
	for i := 2; i < N; i++ {
		B[i] = ni(sc)
	}

	dp := make([]int, N)
	for i := 1; i < N; i++ {
		if i < 2 {
			dp[i] = dp[i-1] + A[i]
			continue
		}

		beforeOne := dp[i-1] + A[i]
		beforeTwo := dp[i-2] + B[i]

		if beforeOne < beforeTwo {
			dp[i] = beforeOne
		} else {
			dp[i] = beforeTwo
		}
	}

	fmt.Fprint(writer, dp[len(dp)-1])
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
