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

	// N=品物の数
	// W=最大重量
	N, W := ni2(sc)

	w, v := make([]int, N+1), make([]int, N+1)
	for i := 1; i < N+1; i++ {
		w[i], v[i] = ni2(sc)
	}

	// 横軸である価値の最大値を計算
	xSize := sumInts(v)

	// dp初期化
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, xSize+1)
	}

	// dp[0]の補完.
	// index=0意外は、全品物の合計重量以上を入れておけばminで比較する時に最小値にはならない.
	// +100はなんとなく
	allSumW := sumInts(w) + 100
	for i := range dp[0] {
		if i == 0 {
			dp[0][0] = 0
		} else {
			dp[0][i] = allSumW
		}
	}

	// i=品物Index
	for i := 1; i < N+1; i++ {
		// j=価値合計Index
		for j := 0; j < xSize+1; j++ {
			targetW := w[i]
			targetV := v[i]

			if j < targetV {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-targetV]+targetW)
			}
		}
	}

	maxV := 0
	for sumV, sumW := range dp[N] {
		if sumW <= W {
			maxV = sumV
		}
	}

	fmt.Fprint(writer, maxV)
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
