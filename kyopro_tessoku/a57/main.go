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

// N Q
// A1 A2 ... An
// X1 Y1
// X2 Y2
// ...
// XQ YQ
func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	// 長い行を読み込めるようにするための定義
	sc.Buffer([]byte{}, MaxInt)

	// N=穴の数
	// Q=質問の数
	N, Q := ni2(sc)

	A := make([]int, N+1)

	for i := 1; i <= N; i++ {
		A[i] = ni(sc)
	}

	// i=0 1日後
	// i=1 2日後
	// i=2 4日後
	// i=3 8日後
	// ...
	// 1≤Y≤10^9 なので 2進数の30桁分あればOK
	dp := make([][]int, 30)
	for i := range dp {
		dp[i] = make([]int, N+1)
	}

	// 穴=1Index始まり
	for i := 1; i <= N; i++ {
		dp[0][i] = A[i]
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j <= N; j++ {
			dp[i][j] = dp[i-1][dp[i-1][j]]
		}
	}

	// Queryに答えていく
	for range make([]struct{}, Q) {
		X, Y := ni2(sc)

		// 2進数に変換
		b := fmt.Sprintf("%b", Y)

		// 2進数の桁数
		l := len(b)

		// 穴の位置
		pos := X

		for i := 0; i < l; i++ {
			n := l - 1 - i
			if b[n] == '1' {
				pos = dp[i][pos]
			}
		}

		fmt.Fprintln(writer, pos)
	}

	// bitが立っている桁の別確認法
	// // Queryに答えていく
	// for range make([]struct{}, Q) {
	// 	X, Y := ni2(sc)

	// 	pos := X
	// 	for d := 30; d >= 0; d-- {
	// 		if Y&(1<<d) != 0 {
	// 			pos = dp[d][pos]
	// 		}
	// 	}

	// 	fmt.Fprintln(writer, pos)
	// }
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

func ns(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
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

// 見つかった場合対象のindexを返す
func Search(s []int, v int) (int, bool) {
	if len(s) <= 0 {
		return -1, false
	}

	return searchHelper(s, v, 0, len(s)-1)
}

func searchHelper(s []int, v, l, r int) (int, bool) {
	if l > r {
		return -1, false
	}

	mid := (l + r) / 2

	if s[mid] == v {
		return mid, true
	}

	if v < s[mid] {
		return searchHelper(s, v, l, mid-1)
	} else {
		return searchHelper(s, v, mid+1, r)
	}
}
