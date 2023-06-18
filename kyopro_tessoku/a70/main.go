package main

import (
	"bufio"
	"container/list"
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

// operation は ランプ z, y, z をスイッチする操作
type operation struct {
	x, y, z int
}

func newOperation(x, y, z int) *operation {
	return &operation{
		x: x,
		y: y,
		z: z,
	}
}

// exec は pos に対して操作を実行した結果の頂点番号を返す
//
// ランプ番号は以下のようになる
// 0 1 0 1 <- 2進数
// 4 3 2 1 <- ランプ番号
func (o *operation) nextVIndex(pos, digits int) int {
	// 2進数へ変換
	//
	// bits = 2進数の桁数の添字にbitを格納したスライス
	// 例)
	// 桁数  0  1  2  3  4
	//     [0, 1, 0, 1, 0]
	// => 0101
	bits := make([]int, digits+1)
	for i := 1; i <= digits; i++ {
		n := 1 << (i - 1)
		bits[i] = (pos / n) % 2
	}

	// ON/OFFのスイッチング
	bits[o.x] = 1 - bits[o.x]
	bits[o.y] = 1 - bits[o.y]
	bits[o.z] = 1 - bits[o.z]

	// 10進数へ変換
	ret := 0
	for i := 1; i < len(bits); i++ {
		bit := bits[i]
		if bit == 1 {
			ret += 1 << (i - 1)
		}
	}

	return ret
}

// getPosition は、引数 s のスライスを 10進数 に変換する
// s は2進数の各桁が添字に格納されているスライス
func getDecimalNumber(s []int) int {
	ret := 0
	for i := 1; i < len(s); i++ {
		if s[i] == 1 {
			ret += 1 << (i - 1)
		}
	}
	return ret
}

type queue struct {
	list *list.List
}

func newQueue() *queue {
	return &queue{
		list: list.New(),
	}
}

func (q *queue) enqueue(n int) {
	q.list.PushBack(n)
}

func (q *queue) dequeue() int {
	if q.list.Len() <= 0 {
		return -1
	}

	el := q.list.Front()
	q.list.Remove(el)
	return el.Value.(int)
}

func (q *queue) isEmpty() bool {
	return q.list.Len() == 0
}

const MaxN = 10

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	// 長い行を読み込めるようにするための定義
	sc.Buffer([]byte{}, MaxInt)

	// N=頂点数
	// M=操作種類数
	N, M := ni2(sc)

	// A=ランプ点灯初期状態
	// 1始まりIndex
	A := make([]int, N+1)
	for i := 1; i < len(A); i++ {
		A[i] = ni(sc)
	}

	operations := make([]*operation, M)
	for i := range operations {
		operations[i] = newOperation(ni3(sc))
	}

	// Graph の作成
	vCount := (1 << N)
	type vIndex int // 頂点番号の型
	graph := make([][]vIndex, vCount)
	for i := 0; i < vCount; i++ {
		for _, ope := range operations {
			next := ope.nextVIndex(i, MaxN+1)
			graph[i] = append(graph[i], vIndex(next))
		}
	}

	// Start/Goal の頂点番号を求める
	start := getDecimalNumber(A)
	goal := (1 << N) - 1

	// 幅優先探索で最短経路を求める
	dists := make([]int, vCount)
	for i := range dists {
		dists[i] = -1 // -1 は未確定頂点の意味
	}
	dists[start] = 0
	q := newQueue()
	q.enqueue(start)

	for !q.isEmpty() {
		from := q.dequeue()

		for _, to := range graph[from] {
			if dists[to] == -1 {
				dists[to] = dists[from] + 1
				q.enqueue(int(to))
			}
		}
	}

	answer := dists[goal]
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
