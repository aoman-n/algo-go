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

type edge struct {
	to  int
	cap int
	rev int
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	// 長い行を読み込めるようにするための定義
	sc.Buffer([]byte{}, MaxInt)

	N, M := ni2(sc)

	// residual graph(残余グラフ)の初期化
	g := make([][]*edge, N+1)
	for range make([]struct{}, M) {
		// a->bへ流量c
		a, b, c := ni3(sc)

		forwardIndex := len(g[a])
		reverseIndex := len(g[b])

		g[a] = append(g[a], &edge{
			to:  b,
			cap: c,
			rev: reverseIndex,
		})
		g[b] = append(g[b], &edge{
			to:  a,
			cap: 0,
			rev: forwardIndex,
		})
	}

	// start->endのルートを深さ優先探索する
	// ルートが見つからなくなった時点で終了
	totalFlow := 0
	for {
		flow := getMinFlow(g, getUsed(N), 1, N, MaxInt)
		if flow == 0 {
			break
		}
		totalFlow += flow
	}

	fmt.Fprint(writer, totalFlow)
}

func getUsed(n int) []bool {
	return make([]bool, n+1)
}

func getMinFlow(g [][]*edge, used []bool, pos, goal, currFlow int) int {
	if pos == goal {
		return currFlow
	}

	used[pos] = true

	for i, edge := range g[pos] {
		if edge.cap == 0 {
			continue
		}
		if used[edge.to] {
			continue
		}

		minFlow := min(currFlow, edge.cap)
		flow := getMinFlow(g, used, edge.to, goal, minFlow)
		if flow >= 1 {
			g[pos][i].cap -= flow
			g[g[pos][i].to][g[pos][i].rev].cap += flow
			return flow
		}
	}

	return 0
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
