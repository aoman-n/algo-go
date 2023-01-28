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

	// N=クーポン枚数
	// M=品物数
	N, M := ni2(sc)

	// [クーポンIndex][品物Index] =
	//   1の場合 => 無料になる品物
	//   0の場合 => 無料にならない品物
	A := make([][]int, M+1)
	for i := range A {
		A[i] = make([]int, N+1)
		if i == 0 {
			continue
		}

		for j := 1; j < len(A[i]); j++ {
			A[i][j] = ni(sc)
		}
	}

	var MaximumNum = M + 100

	// dpの作成
	// dp[クーポンIndex][無料になる品物の集合]
	// 集合はbitで処理し10進数で表現
	dp := make([][]int, M+1)
	for i := range dp {
		dp[i] = make([]int, 1<<N)
		for j := range dp[i] {
			dp[i][j] = MaximumNum
		}
	}

	// dpの初期化
	for i := range dp {
		dp[i][0] = 0
	}

	// i=クーポンIndex
	// j=品物の集合Index
	for i := 1; i < len(dp); i++ {
		for j := 0; j < 1<<N; j++ {
			currS := 0
			// k=品物Index
			// v=無料になるかどうか(0or1)
			for k, v := range A[i] {
				if v != 0 {
					currS += 1 << (k - 1)
				}
			}

			selectedPatternS := currS | j

			dp[i][j] = min(dp[i-1][j], dp[i][j])
			dp[i][selectedPatternS] = min(dp[i][selectedPatternS], dp[i][j]+1)
		}
	}

	answer := dp[M][(1<<N)-1]
	if answer == MaximumNum {
		fmt.Fprint(writer, -1)
	} else {
		fmt.Fprint(writer, answer)
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
