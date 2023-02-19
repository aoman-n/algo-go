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

	// N=品物数
	// M=クーポン枚数
	N, M := ni2(sc)

	A := make([][]int, M+1)
	for i := 1; i < len(A); i++ {
		row := make([]int, N+1)
		for j := 1; j < len(row); j++ {
			row[j] = ni(sc)
		}
		A[i] = row
	}

	// maxXIndex = 全部の品物を選ぶ場合の10進数
	maxXIndex := (1 << N) - 1

	dp := make([][]int, M+1)
	for i := range dp {
		dp[i] = make([]int, maxXIndex+1)
	}

	// dpの初期化
	outOfCouponCount := M + 1000 // 1000は適当な数値
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = outOfCouponCount
		}
	}
	for i := range dp {
		dp[i][0] = 0
	}

	// i=クーポンIndex
	// j=品物の集合(S)Index
	for i := 1; i < M+1; i++ {
		for j := 0; j < 1<<N; j++ {
			// currCouponS = 現在の対象クーポンが無料になる品物の集合(S)の10進数
			currCouponS := 0
			// k = 品物Index
			for k, v := range A[i] {
				// kが無料になるクーポンであれば、currCouponS(集合)に加える
				if v == 1 {
					currCouponS += 1 << (k - 1)
				}
			}

			// selectedCouponS = 現在のクーポンを選んだ場合の集合
			selectedCouponS := currCouponS | j

			dp[i][j] = min(dp[i][j], dp[i-1][j])
			dp[i][selectedCouponS] = min(dp[i][j]+1, dp[i][selectedCouponS])
		}
	}

	answer := dp[M][(1<<N)-1]
	if answer == outOfCouponCount {
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
