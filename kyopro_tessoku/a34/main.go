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
	N, X, Y := ni3(sc)
	A := nis(sc, N)

	maxM := -1
	for _, v := range A {
		if v > maxM {
			maxM = v
		}
	}

	dp := grundyDP(maxM, X, Y)

	grundies := make([]int, 0, N)
	for _, v := range A {
		grundies = append(grundies, dp[v])
	}

	sumXor := grundies[0]
	for i := 1; i < len(grundies); i++ {
		sumXor = sumXor ^ grundies[i]
	}

	if sumXor == 0 {
		fmt.Fprint(writer, "Second")
	} else {
		fmt.Fprint(writer, "First")
	}

}

func grundyDP(m, x, y int) []int {
	dp := make([]int, m+1)

	for i := 1; i < len(dp); i++ {
		a := make([]int, 0, 3)

		if i >= x {
			a = append(a, dp[i-x])
		}
		if i >= y {
			a = append(a, dp[i-y])
		}

		if len(a) > 0 {
			dp[i] = get(a)
		} else {
			dp[i] = 0
		}
	}

	return dp
}

// 0,1,2の存在しない最小値を返す
func get(a []int) int {
	m := map[int]bool{
		0: true,
		1: true,
		2: true,
	}

	for _, v := range a {
		m[v] = false
	}

	ret := 100
	for v, flag := range m {
		if flag && v < ret {
			ret = v
		}
	}
	return ret
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
