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
	N := ni(sc)

	P := make([]int, N+2)
	A := make([]int, N+2)
	for i := 1; i < N+1; i++ {
		P[i], A[i] = ni2(sc)
	}

	// 処理を簡易化して、indexアクセスエラーにならないように+2している
	dp := make([][]int, N+2)
	for i := range dp {
		dp[i] = make([]int, N+2)
	}

	// l=縦列(left)
	// r=横列(right)
	// 区間(l, r)
	// l=rになったらすべて取り除かれている状態なので終了
	for l := 1; l < N+1; l++ {
		for r := N; r >= 0; r-- {
			if l > r {
				break
			}

			originalR := r + 1
			fromRightPoint := dp[l][originalR]
			// (l<=p<=r)であれば得点
			// r+1を削除するのでP[r+1]を取得
			if l <= P[originalR] && P[originalR] <= originalR {
				fromRightPoint += A[originalR]
			}

			originalL := l - 1
			fromLeftPoint := dp[originalL][r]
			// (l<=p<=r)であれば得点
			// l-1を削除するのでP[l-1]を取得
			if originalL <= P[originalL] && P[originalL] <= r {
				fromLeftPoint += A[originalL]
			}

			dp[l][r] = maxInt(fromLeftPoint, fromRightPoint)
		}
	}

	answer := 0
	for i := range dp {
		answer = maxInt(answer, dp[i][i])
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

func maxInt(s ...int) int {
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
