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

	sc.Scan()
	S := sc.Text()
	sc.Scan()
	T := sc.Text()

	dp := make([][]int, len(S)+1)
	for i := range dp {
		dp[i] = make([]int, len(T)+1)
	}

	// dpの初期化
	// 1行目と1列目の補完
	for i := range dp[0] {
		dp[0][i] = i
	}
	for i := range dp {
		dp[i][0] = i
	}

	// ------ dp二次元配列の遷移の操作 --------
	// ① ↓ : Sへの挿入 (コスト1)
	// ② → : Sの削除 (コスト1)
	// ③ ➘ : Sの変更 (コスト0 or 1)
	//
	// ③の操作のみコスト0の可能性がある。なぜなら文字の長さが同じであるため。
	// 対象座標の文字列が一致した場合のみコスト0となる。
	// -------------------------------------

	// i=S文字列のIndex
	// j=T文字列のIndex
	for i := 1; i < len(S)+1; i++ {
		for j := 1; j < len(T)+1; j++ {
			currS := S[i-1]
			currT := T[j-1]

			changedOperationCost := dp[i-1][j-1]
			if currS != currT {
				changedOperationCost++
			}

			dp[i][j] = minInt(
				// ①↓の遷移
				dp[i-1][j]+1,
				// ②→の遷移
				dp[i][j-1]+1,
				// ③➘の遷移
				changedOperationCost,
			)
		}
	}

	fmt.Fprint(writer, dp[len(S)][len(T)])
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

func minInt(s ...int) int {
	if len(s) <= 0 {
		return MinInt
	}

	min := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
	}

	return min
}
