package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
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

	C := make([][]string, H)
	for i := range C {
		rowStr := nt(sc)
		C[i] = strings.Split(rowStr, "")
	}

	dp := make([][]int, H)
	for i := range dp {
		dp[i] = make([]int, W)
	}

	dp[0][0] = 1

	const (
		WriteBlock = "."
		BlackBlock = "#"
	)

	// y=縦のIndex
	// x=横のIndex
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if y == 0 && x == 0 {
				continue
			}

			currBlock := C[y][x]

			if currBlock == BlackBlock {
				// 黒いブロックの場合は行く方法がないので0になる
				dp[y][x] = 0
			} else {
				if y == 0 {
					dp[y][x] = dp[y][x-1]
				} else if x == 0 {
					dp[y][x] = dp[y-1][x]
				} else {
					dp[y][x] = dp[y-1][x] + dp[y][x-1]
				}
			}
		}
	}

	answer := dp[H-1][W-1]
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

func nt(sc *bufio.Scanner) string {
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
