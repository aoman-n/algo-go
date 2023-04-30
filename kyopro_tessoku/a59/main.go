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

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

type rsq struct {
	data       []int
	bottomSize int
	treeSize   int
}

func newRsq(n int) *rsq {
	// bottomSizeの計算
	x := 0
	for {
		if pow(2, x) >= n {
			break
		}
		x++
	}
	bottomSize := pow(2, x)
	treeSize := bottomSize*2 - 1

	return &rsq{
		// 1Indexで計算処理したいので+1する
		data:       make([]int, treeSize+1),
		bottomSize: bottomSize,
		treeSize:   treeSize,
	}
}

func (rs *rsq) update(pos, x int) {
	bottomPos := rs.bottomSize + pos - 1
	rs.data[bottomPos] = x

	curr := rs.parent(bottomPos)
	for curr >= 1 {
		lPos := rs.left(curr)
		rPos := rs.right(curr)
		rs.data[curr] = rs.data[lPos] + rs.data[rPos]

		curr = rs.parent(curr)
	}
}

func (rs *rsq) parent(i int) int {
	return i / 2
}

func (rs *rsq) left(i int) int {
	return i * 2
}

func (rs *rsq) right(i int) int {
	return i*2 + 1
}

func (rs *rsq) getSumInRange(l, r int) int {
	return rs.getSumInRangeHelper(l, r, 1, rs.bottomSize+1, 1)
}

func (rs *rsq) getSumInRangeHelper(targetL, targetR, currL, currR, currPos int) int {
	// 区間外
	if currL >= targetR || currR <= targetL {
		return 0
	}
	// 区間内
	if currL >= targetL && currR <= targetR {
		return rs.data[currPos]
	}

	mid := (currL + currR) / 2
	lSum := rs.getSumInRangeHelper(targetL, targetR, currL, mid, rs.left(currPos))
	rSum := rs.getSumInRangeHelper(targetL, targetR, mid, currR, rs.right(currPos))
	return lSum + rSum
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	// 長い行を読み込めるようにするための定義
	sc.Buffer([]byte{}, MaxInt)

	N, Q := ni2(sc)

	rsq := newRsq(N)

	const (
		UpdateQuery = 1
		AnswerQuery = 2
	)

	for range make([]struct{}, Q) {
		query := ni(sc)

		switch query {
		case UpdateQuery:
			pos, x := ni2(sc)
			rsq.update(pos, x)
		case AnswerQuery:
			l, r := ni2(sc)
			sum := rsq.getSumInRange(l, r)
			fmt.Fprintln(writer, sum)
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
