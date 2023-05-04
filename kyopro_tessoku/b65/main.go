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
	// 長い行を読み込めるようにするための定義
	sc.Buffer([]byte{}, MaxInt)

	// N:社員数
	// T:社長
	N, T := ni2(sc)

	g := make(map[int][]int)

	for range make([]struct{}, N-1) {
		a, b := ni2(sc)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	dp := make([]int, N+1)
	visited := make([]bool, N+1)
	exec(dp, g, visited, T)

	for i := 1; i < len(dp); i++ {
		if i == 1 {
			fmt.Fprintf(writer, "%d", dp[i])
		} else {
			fmt.Fprintf(writer, " %d", dp[i])
		}
	}
	fmt.Fprintln(writer)
}

func exec(dp []int, g map[int][]int, visited []bool, curr int) {
	visited[curr] = true
	children := excludeVisited(visited, g[curr])
	if len(children) == 0 {
		dp[curr] = 0
		return
	}

	for _, child := range children {
		exec(dp, g, visited, child)
	}

	max := -1
	for _, child := range children {
		if dp[child] > max {
			max = dp[child]
		}
	}

	dp[curr] = max + 1
}

func excludeVisited(visited []bool, s []int) []int {
	ret := make([]int, 0, len(s))
	for _, v := range s {
		if !visited[v] {
			ret = append(ret, v)
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
