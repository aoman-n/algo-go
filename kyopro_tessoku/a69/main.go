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

func newEdge(to, cap, rev int) *edge {
	return &edge{
		to:  to,
		cap: cap,
		rev: rev,
	}
}

type MaxFlow struct {
	graph [][]*edge
	used  []bool
}

func NewMaxFlow(n int) *MaxFlow {
	return &MaxFlow{
		graph: make([][]*edge, n+1),
		used:  make([]bool, n+1),
	}
}

func (m *MaxFlow) resetUsed() {
	for i := range m.used {
		m.used[i] = false
	}
}

func (m *MaxFlow) AddEdge(a, b, cap int) {
	currALen := len(m.graph[a])
	currBLen := len(m.graph[b])
	m.graph[a] = append(m.graph[a], newEdge(b, cap, currBLen))
	m.graph[b] = append(m.graph[b], newEdge(a, 0, currALen))
}

func (m *MaxFlow) findFlow(pos, goal, f int) int {
	if pos == goal {
		return f
	}

	m.used[pos] = true

	for _, edge := range m.graph[pos] {
		if m.used[edge.to] {
			continue
		}
		if edge.cap <= 0 {
			continue
		}

		flow := m.findFlow(edge.to, goal, min(f, edge.cap))

		if flow >= 1 {
			edge.cap -= flow
			m.graph[edge.to][edge.rev].cap += flow
			return flow
		}
	}

	return 0
}

func (m *MaxFlow) Get(s, t int) int {
	totalFlow := 0

	for {
		m.resetUsed()

		flow := m.findFlow(s, t, maxVertexLen+200) // 200はなんとなく

		if flow <= 0 {
			break
		}

		totalFlow += flow
	}

	return totalFlow
}

// 1≤N≤150
const maxVertexLen = (150 * 2) + 2

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	// 長い行を読み込めるようにするための定義
	sc.Buffer([]byte{}, MaxInt)

	// N=生徒数
	N := ni(sc)

	const (
		ok = "#"
		ng = "."
	)

	// <すべての頂点を数値に変換する>
	var (
		// 生徒の頂点番号: 生徒番号そのまま
		// 席の頂点番号: N + 席番号
		// スタートの頂点: (N * 2) + 1
		// エンドの頂点: (N * 2) + 2
		startVertex  = (N * 2) + 1
		endVertex    = (N * 2) + 2
		allVertexLen = (N * 2) + 2
	)

	maxFlow := NewMaxFlow(allVertexLen)

	for i := range make([]struct{}, N) {
		studentNum := i + 1
		str := ns(sc)

		for i, r := range str {
			if string(r) == ok {
				seatNum := N + (i + 1)
				maxFlow.AddEdge(studentNum, seatNum, 1)
			}
		}
	}

	// スタート頂点 -> 生徒頂点 を結ぶ
	// 座頂点 -> エンド頂点 を結ぶ
	for i := 1; i <= N; i++ {
		studentNum := i
		seatNum := N + i
		maxFlow.AddEdge(startVertex, studentNum, 1)
		maxFlow.AddEdge(seatNum, endVertex, 1)
	}

	answer := maxFlow.Get(startVertex, endVertex)
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
