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
	parents map[int]int
	// levels: root頂点から末端までの最大距離
	levels map[int]int
}

// n:作成時の頂点の数
func newUnionFind(n int) *unionFind {
	parents := make(map[int]int)
	levels := make(map[int]int)
	for i := 1; i <= n; i++ {
		parents[i] = i
		levels[i] = 1
	}

	return &unionFind{
		parents: parents,
		levels:  levels,
	}
}

// unite: 頂点uと頂点vを結合する
func (uf *unionFind) unite(u, v int) {
	rootU, rootV := uf.root(u), uf.root(v)
	if rootU == rootV {
		return
	}

	// levelが大きい方に結合する
	if uf.levels[rootU] >= uf.levels[rootV] {
		uf.parents[rootV] = rootU
	} else {
		uf.parents[rootU] = rootV
	}

	// levelが同じ頂点同士を結合したときのみ、levelが1増加する
	if uf.levels[rootU] == uf.levels[rootV] {
		uf.levels[rootU]++
	}
}

// root: xのrootを返す。
func (uf *unionFind) root(x int) int {
	ret := x
	for {
		if uf.parents[ret] == ret {
			break
		}
		ret = uf.parents[ret]
	}

	return ret
}

// same: 頂点uと頂点vが同じグループに属しているかを返す
func (uf *unionFind) same(u, v int) bool {
	return uf.root(u) == uf.root(v)
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	// 長い行を読み込めるようにするための定義
	sc.Buffer([]byte{}, MaxInt)
	N, Q := ni2(sc)

	const (
		queryUnite  = 1
		queryAnswer = 2
	)

	unionFind := newUnionFind(N)

	for range make([]struct{}, Q) {
		query, u, v := ni3(sc)

		switch query {
		case queryUnite:
			unionFind.unite(u, v)
		case queryAnswer:
			if unionFind.same(u, v) {
				fmt.Fprintln(writer, "Yes")
			} else {
				fmt.Fprintln(writer, "No")
			}
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
