package main

import (
	"bufio"
	"container/heap"
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

type relation struct {
	/* 向かい先の頂点番号 */
	to int
	/* toへの距離 */
	dist int
}

func newRelation(to, dist int) *relation {
	return &relation{
		to:   to,
		dist: dist,
	}
}

type vertex struct {
	/* 頂点識別子番号 */
	number int
	/* 最短経路の距離 */
	currMinDist int
}

type priorityQueue []*vertex

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].currMinDist > pq[j].currMinDist
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	vertex := x.(*vertex)
	*pq = append(*pq, vertex)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	vertex := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return vertex
}

type vertexPriorityQueue struct {
	data priorityQueue
}

func newVertexPriorityQueue() *vertexPriorityQueue {
	vertexList := make(priorityQueue, 0)
	heap.Init(&vertexList)
	return &vertexPriorityQueue{
		data: vertexList,
	}
}

func (pq *vertexPriorityQueue) push(v *vertex) {
	heap.Push(&pq.data, v)
}

func (pq *vertexPriorityQueue) pop() *vertex {
	v := heap.Pop(&pq.data)
	if v == nil {
		return nil
	}

	return v.(*vertex)
}

func (pq *vertexPriorityQueue) len() int {
	return pq.data.Len()
}

// TODO: 一部3秒以上かかってfailになるのを修正
func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	// 長い行を読み込めるようにするための定義
	sc.Buffer([]byte{}, MaxInt)

	N, M := ni2(sc)

	relationMap := make(map[int][]*relation)
	for range make([]struct{}, M) {
		x, y, dist := ni3(sc)
		relationMap[x] = append(relationMap[x], newRelation(y, dist))
		relationMap[y] = append(relationMap[y], newRelation(x, dist))
	}

	// 頂点1からスタート
	startVertex := &vertex{
		number:      1,
		currMinDist: 0,
	}
	pq := newVertexPriorityQueue()
	pq.push(startVertex)

	// dists: 現状の最短経路距離を保持するリスト。Indexが頂点番号
	dists := make([]int, N+1)
	for i := range dists {
		dists[i] = MaxInt
	}
	dists[1] = 0

	for pq.len() > 0 {
		v := pq.pop()
		relations := relationMap[v.number]

		for _, r := range relations {
			dist := r.dist + v.currMinDist

			// すでに入っている距離が本経路よりも最短であれば処理しない
			if dists[r.to] < dist {
				continue
			}

			dists[r.to] = dist
			pq.push(&vertex{
				number:      r.to,
				currMinDist: dist,
			})
		}
	}

	for i := 1; i < len(dists); i++ {
		if dists[i] == MaxInt {
			fmt.Fprintln(writer, -1)
		} else {
			fmt.Fprintln(writer, dists[i])
		}
	}
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
