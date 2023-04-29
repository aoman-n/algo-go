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

type rsq struct {
	nodes     []int
	bottomLen int
}

func newRsq(size int) *rsq {
	bottomLen := getBottomLen(size)
	nodeLen := bottomLen*2 - 1
	// 1始まりのIndexでアクセスしたいので+1したtreeのnode数にする
	treeSize := nodeLen + 1

	return &rsq{
		// 各nodeは0で初期化
		nodes:     make([]int, treeSize),
		bottomLen: bottomLen,
	}
}

func getBottomLen(nodeLen int) int {
	n := 0
	for pow(2, n) < nodeLen {
		n++
	}

	return pow(2, n)
}

// pos: 更新するposition
// x  : 更新する値
func (r *rsq) update(pos, x int) {
	bottomPos := pos
	treePos := r.bottomLen + bottomPos - 1

	r.nodes[treePos] = x

	currPos := r.parent(treePos)
	for currPos > 0 {
		// leftとrightのmax値に更新
		leftVal := r.nodes[r.left(currPos)]
		rightVal := r.nodes[r.right(currPos)]
		r.nodes[currPos] = max(leftVal, rightVal)
		// 親のnodeへ移動
		currPos = r.parent(currPos)
	}
}

func (r *rsq) parent(i int) int {
	return i / 2
}

func (r *rsq) left(i int) int {
	return i * 2
}

func (r *rsq) right(i int) int {
	return i*2 + 1
}

// [l,r)半開区間の最大値を返す
func (rs *rsq) getMaxInRange(l, r int) int {
	return rs.getMaxInRangeHelper(l, r, 1, rs.bottomLen+1, 1)
}

func (rs *rsq) getMaxInRangeHelper(targetL, targetR, currL, currR, currPos int) int {
	// 範囲外の場合は0を返す
	if currL > targetR || currR < targetL {
		return 0
	}
	// 範囲に完全に入っている場合はcurrPosの値を返す
	if currL >= targetL && currR <= targetR {
		return rs.nodes[currPos]
	}
	// leafNodeで範囲外ならreturn
	if currR-currL <= 1 {
		return 0
	}

	mid := (currL + currR) / 2
	lMax := rs.getMaxInRangeHelper(targetL, targetR, currL, mid, rs.left(currPos))
	rMax := rs.getMaxInRangeHelper(targetL, targetR, mid, currR, rs.right(currPos))
	return max(lMax, rMax)
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	// 長い行を読み込めるようにするための定義
	sc.Buffer([]byte{}, MaxInt)

	N, Q := ni2(sc)

	rsq := newRsq(N)

	const (
		updateQ        = 1
		getRangeInMaxQ = 2
	)

	for range make([]struct{}, Q) {
		q := ni(sc)

		switch q {
		case updateQ:
			pos, x := ni2(sc)
			rsq.update(pos, x)
		case getRangeInMaxQ:
			l, r := ni2(sc)
			maxInRange := rsq.getMaxInRange(l, r)
			fmt.Fprintln(writer, maxInRange)
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

// a=底
// b=指数
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
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
