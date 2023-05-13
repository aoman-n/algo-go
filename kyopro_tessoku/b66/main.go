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

type unionFind struct {
	parents []int
	depths  []int
}

func newUnionFind(n int) *unionFind {
	parents := make([]int, n+1)
	depths := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parents[i] = i
		depths[i] = 1
	}

	return &unionFind{
		parents: parents,
		depths:  depths,
	}
}

// 結合
func (uf *unionFind) unite(a, b int) {
	rootA := uf.root(a)
	rootB := uf.root(b)

	if rootA == rootB {
		return
	}

	if uf.depths[rootA] >= uf.depths[rootB] {
		uf.parents[rootB] = rootA
	} else {
		uf.parents[rootA] = rootB
	}

	if uf.depths[rootA] == uf.depths[rootB] {
		uf.depths[rootA]++
	}
}

func (uf *unionFind) root(i int) int {
	ret := i
	for {
		if ret == uf.parents[ret] {
			break
		}
		ret = uf.parents[ret]
	}

	return ret
}

func (uf *unionFind) same(a, b int) bool {
	return uf.root(a) == uf.root(b)
}

type route struct {
	a, b    int
	enabled bool
}

func newRoute(a, b int) *route {
	return &route{
		a:       a,
		b:       b,
		enabled: true,
	}
}

func (r *route) disabled() {
	r.enabled = false
}

// x本目の路線を運休にするquery
type query1 struct {
	x int
}

func newQuery1(x int) *query1 {
	return &query1{
		x: x,
	}
}

func (q *query1) isQuery() {}

// aとbの駅が繋がっているか答えるquery
type query2 struct {
	a, b int
}

func newQuery2(a, b int) *query2 {
	return &query2{
		a: a,
		b: b,
	}
}

func (q *query2) isQuery() {}

type query interface {
	isQuery()
}

type queries []query

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	// 長い行を読み込めるようにするための定義
	sc.Buffer([]byte{}, MaxInt)

	N, M := ni2(sc)

	routes := make([]*route, M+1)
	for i := range make([]struct{}, M) {
		a, b := ni2(sc)
		routes[i+1] = newRoute(a, b)
	}

	Q := ni(sc)
	queries := make(queries, 0, Q)
	for range make([]struct{}, Q) {
		q := ni(sc)

		switch q {
		case 1: // x本目の路線を運休させるクエリ
			x := ni(sc)
			queries = append(queries, newQuery1(x))
			routes[x].disabled()
		case 2: // aとbの駅がつながっているか解答するクエリ
			a, b := ni2(sc)
			queries = append(queries, newQuery2(a, b))
		}
	}

	uf := newUnionFind(N)

	for i := 1; i < len(routes); i++ {
		route := routes[i]
		if route.enabled {
			uf.unite(route.a, route.b)
		}
	}

	type answerString string

	const (
		answerYes answerString = "Yes"
		answerNo  answerString = "No"
	)

	answers := make([]answerString, 0, len(queries))

	for i := len(queries) - 1; i >= 0; i-- {
		query := queries[i]

		switch q := query.(type) {
		case *query1:
			route := routes[q.x]
			uf.unite(route.a, route.b)
		case *query2:
			if uf.same(q.a, q.b) {
				answers = append(answers, answerYes)
			} else {
				answers = append(answers, answerNo)
			}
		}
	}

	for i := len(answers) - 1; i >= 0; i-- {
		fmt.Fprintln(writer, answers[i])
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
