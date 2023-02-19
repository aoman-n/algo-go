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

	// N=ブロックの数
	N := ni(sc)

	// P=残っていたらポイントとなるブロックIndex
	// A=獲得ポイント
	P, A := make([]int, N+2), make([]int, N+2)
	for i := 1; i < N+1; i++ {
		P[i], A[i] = ni2(sc)
	}

	dp := make([][]int, N+2)
	for i := range dp {
		// マス目の "<-" 移動があるため、処理しやすいように1つ伸ばしておく
		// +1 は1Indexの対応
		// +1 は"<-"移動の対応
		dp[i] = make([]int, N+2)
	}

	/**
	(l,r)
	- l=left のposition
	- r=right のposition

	(l,r)=(1,N) のポジションからスタートさせる
	**/
	for l := 1; l <= N; l++ {
		for r := N; r >= 1; r-- {
			if l > r {
				break
			}

			// right からpop
			popRIndex := r + 1
			pointFromR := dp[l][r+1]
			// point 獲得判定
			// P[popRIndex] のブロックが残っていれば point 獲得!!
			if l <= P[popRIndex] && P[popRIndex] <= r {
				pointFromR += A[popRIndex]
			}

			// left からpop
			popLIndex := l - 1
			pointFromL := dp[l-1][r]
			// point 獲得判定
			// P[popLIndex] のブロックが残っていれば point 獲得!!
			if l <= P[popLIndex] && P[popLIndex] <= r {
				pointFromL += A[popLIndex]
			}

			dp[l][r] = max(pointFromL, pointFromR)
		}
	}

	answer := 0
	for i := 1; i <= N; i++ {
		answer = max(answer, dp[i][i])
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
