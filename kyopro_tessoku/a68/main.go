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
	// to: 行き先の頂点
	to int
	// cap: to への残余
	cap int
	// rev: to から自身への edge 保持する index number ( graph[to][x] の x )
	rev int
}

func newEdge(to, cap, rev int) *edge {
	return &edge{
		to:  to,
		cap: cap,
		rev: rev,
	}
}

type MaxFlow struct {
	// graph: 残余グラフ情報を保持する2次元配列
	graph [][]*edge
	// used: 頂点が使用済かを保持する。indexが頂点番号
	used []bool
}

func NewMaxFlow(n int) *MaxFlow {
	return &MaxFlow{
		graph: make([][]*edge, n+1),
		used:  make([]bool, n+1),
	}
}

// AddEdge は 頂点 a と 頂点 b の辺を追加する
func (m *MaxFlow) AddEdge(a, b, cap int) {
	currALen := len(m.graph[a])
	currBLen := len(m.graph[b])
	m.graph[a] = append(m.graph[a], newEdge(b, cap, currBLen))
	m.graph[b] = append(m.graph[b], newEdge(a, 0, currALen))

}

func (m *MaxFlow) resetUsed() {
	for i := range m.used {
		m.used[i] = false
	}
}

// findFlow は 頂点 pos から 頂点 goal までの深さ優先探索を行い、流すことが可能な最大流量を返す。
// 引数 f は現状の最大流量で、再帰的に呼び出すために渡している。
func (m *MaxFlow) findFlow(pos, goal, currMaxFlow int) int {
	if pos == goal {
		return currMaxFlow
	}

	m.used[pos] = true

	for _, edge := range m.graph[pos] {
		// 訪問済みの頂点はスキップ
		if m.used[edge.to] {
			continue
		}
		// 0以下だと水は流せないのでスキップ
		if edge.cap <= 0 {
			continue
		}

		// 現在の pos 以降の最大流量を取得する
		minFlow := m.findFlow(edge.to, goal, min(currMaxFlow, edge.cap))

		if minFlow >= 1 {
			// 残余グラフを更新する
			edge.cap -= minFlow
			m.graph[edge.to][edge.rev].cap += minFlow

			return minFlow
		}
	}

	return 0
}

// Get は 頂点 s から 頂点 t への最大流量を返す
func (m *MaxFlow) Get(s, t int) int {
	totalFlow := 0

	for {
		m.resetUsed()

		// 頂点 s -> t の流すことのできる流量 flow を探索する
		flow := m.findFlow(s, t, maxC+100)
		if flow <= 0 {
			break
		}

		totalFlow += flow
	}

	return totalFlow
}

const (
	maxC = 5000
)

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	// 長い行を読み込めるようにするための定義
	sc.Buffer([]byte{}, MaxInt)

	// N=頂点の数(タンク個数)
	// M=辺の数(パイプの本数)
	N, M := ni2(sc)

	maxFlow := NewMaxFlow(N)

	for range make([]struct{}, M) {
		a, b, cap := ni3(sc)
		maxFlow.AddEdge(a, b, cap)
	}

	answer := maxFlow.Get(1, N)
	fmt.Fprintln(writer, answer)
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
