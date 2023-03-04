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
	H, W := ni2(sc)

	n := H + W - 2
	r := W - 1

	answer := c(n, r)
	fmt.Fprint(writer, answer)
}

func c(n, r int) int {
	top := 1
	for i := 2; i <= n; i++ {
		top = (top * i) % MOD
	}

	bottomL := 1
	for i := 2; i <= r; i++ {
		bottomL = (bottomL * i) % MOD
	}
	bottomR := 1
	for i := 2; i <= (n - r); i++ {
		bottomR = (bottomR * i) % MOD
	}

	bottom := (bottomL * bottomR) % MOD

	return division(top, bottom)
}

func division(a, b int) int {
	return (a * power(b, MOD-2)) % MOD
}

func power(a, b int) int {
	ret := 1
	p := a

	for i := 0; i < 30; i++ {
		if (1<<i)&b != 0 {
			ret = (ret * p) % MOD
		}
		p = (p * p) % MOD
	}

	return ret
}

// 動的計画法だと間に合わない
// func dynamic(H, W int) int {
// 	dp := make([][]int, H+1)
// 	for i := range dp {
// 		dp[i] = make([]int, W+1)
// 	}

// 	dp[1][1] = 1
// 	// i=縦index
// 	// j=横index
// 	for i := 1; i < H+1; i++ {
// 		for j := 1; j < W+1; j++ {
// 			if i == 1 || j == 1 {
// 				dp[i][j] = 1
// 			} else {
// 				dp[i][j] = (dp[i-1][j] + dp[i][j-1]) % MOD
// 			}
// 		}
// 	}

// 	return dp[H][W]
// }

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
