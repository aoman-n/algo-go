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

	// N=カード枚数
	// S=合計値
	N, S := ni2(sc)
	// A=カード
	A := make([]int, N+1)
	for i := 1; i < N+1; i++ {
		A[i] = ni(sc)
	}

	dp := make([][]bool, N+1)
	for i := range dp {
		dp[i] = make([]bool, S+1)
	}

	dp[0][0] = true

	// i=カードIndex
	// j=合計値Index
	for i := 1; i < N+1; i++ {
		for j := 0; j < S+1; j++ {
			cardNum := A[i]

			if j >= cardNum {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-cardNum]
			} else {
				dp[i][j] = dp[i-1][j]
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

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func sumInts(s []int) int {
	var ret int
	for _, v := range s {
		ret += v
	}
	return ret
}

func maxInt(s []int) int {
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
