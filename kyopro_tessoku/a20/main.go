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
	S := niT(sc)
	T := niT(sc)

	sAry := make([]string, len([]rune(S))+1)
	for i, r := range S {
		sAry[i+1] = string(r)
	}

	tAry := make([]string, len([]rune(T))+1)
	for i, r := range T {
		tAry[i+1] = string(r)
	}

	// 横軸=S
	// 縦軸=T
	dp := make([][]int, len(tAry))
	for i := range dp {
		dp[i] = make([]int, len(sAry))
	}

	// i=縦軸のindex
	// j=横軸のindex
	for i := 1; i < len(tAry); i++ {
		for j := 1; j < len(sAry); j++ {
			currT := tAry[i]
			currS := sAry[j]

			if currT == currS {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]+1)
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	answer := dp[len(tAry)-1][len(sAry)-1]
	fmt.Fprint(writer, answer)
}

// ==================================================
// io
// ==================================================

func niT(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

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

func max(s ...int) int {
	ret := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > ret {
			ret = s[i]
		}
	}
	return ret
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
