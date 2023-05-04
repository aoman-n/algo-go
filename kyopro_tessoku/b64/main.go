package main

import (
	"bufio"
	"container/heap"
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

type side struct {
	/* 向かう頂点番号 */
	to int
	/* 頂点までの距離 */
	dist int
}

// [][頂点番号,最短経路距離]
type priorityQueue [][2]int

func newPriorityQueue() *priorityQueue {
	return &priorityQueue{}
}

func (h priorityQueue) Len() int           { return len(h) }
func (h priorityQueue) Less(i, j int) bool { return h[i][1] > h[i][1] }
func (h priorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *priorityQueue) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *priorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// TODO:
// 一部実行時間が間に合わない
func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	// 長い行を読み込めるようにするための定義
	sc.Buffer([]byte{}, MaxInt)

	N, M := ni2(sc)

	g := make(map[int][]*side)

	for range make([]struct{}, M) {
		a, b, dist := ni3(sc)
		g[a] = append(g[a], &side{
			to:   b,
			dist: dist,
		})
		g[b] = append(g[b], &side{
			to:   a,
			dist: dist,
		})
	}

	pq := newPriorityQueue()
	heap.Init(pq)
	// 頂点1を入れておく
	heap.Push(pq, [2]int{1, 0})

	// distList: 各頂点への最短距離計算結果を保持するリスト
	resultDists := make([]int, N+1)
	for i := range resultDists {
		resultDists[i] = MaxInt
	}
	resultDists[1] = 0

	for pq.Len() > 0 {
		v := heap.Pop(pq).([2]int)

		sides := g[v[0]]

		for _, side := range sides {
			currDist := side.dist + v[1]

			if resultDists[side.to] > currDist {
				resultDists[side.to] = currDist

				heap.Push(pq, [2]int{side.to, currDist})
			}
		}
	}

	answerRoute := []string{strconv.Itoa(N)}

	currVertexNum := N

	for currVertexNum != 1 {
		dist := resultDists[currVertexNum]
		sides := g[currVertexNum]

		for _, side := range sides {
			if resultDists[side.to] == dist-side.dist {
				// 最短経路で通った辺
				answerRoute = append(answerRoute, strconv.Itoa(side.to))
				currVertexNum = side.to
				break
			}
		}
	}

	// reverse処理
	for i := len(answerRoute)/2 - 1; i >= 0; i-- {
		opp := len(answerRoute) - 1 - i
		answerRoute[i], answerRoute[opp] = answerRoute[opp], answerRoute[i]
	}

	fmt.Fprintln(writer, strings.Join(answerRoute, " "))
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
