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

const (
	MOD = 1000000007
	// ↓AtCoderがGo1.14.1を使用しているため、定義しておく
	intSize = 32 << (^uint(0) >> 63) // 32 or 64
	MaxInt  = 1<<(intSize-1) - 1
	MinInt  = -1 << (intSize - 1)
)

func init() {
	reader = os.Stdin
	writer = os.Stdout
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	// n=品物数
	// w=最大重量
	n, w := ni2(sc)

	W := make([]int, n+1)
	V := make([]int, n+1)

	for i := 1; i < n+1; i++ {
		W[i], V[i] = ni2(sc)
	}

	createRow := func() []int {
		ret := make([]int, w+1)
		for i := range ret {
			ret[i] = MinInt
		}
		return ret
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = createRow()
	}

	dp[0][0] = 0

	// i=品物Index
	// j=合計重量Index
	for i := 1; i < n+1; i++ {
		for j := 0; j < w+1; j++ {
			beforeTotalV := dp[i-1][j]

			currW, currV := W[i], V[i]

			if j < currW {
				dp[i][j] = beforeTotalV
			} else {
				dp[i][j] = max(dp[i-1][j-currW]+currV, beforeTotalV)
			}
		}
	}

	answer := maxIn(dp[n])
	fmt.Fprint(writer, answer)
}

func maxIn(s []int) int {
	if len(s) <= 0 {
		return MinInt
	}

	max := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}

	return max
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

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
