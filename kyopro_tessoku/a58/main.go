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
	data        []int
	leafNodeLen int
	size        int
}

func newRsq(size int) *rsq {
	// sの要素がleafNodeの数に最小で収まるようなleafNodeの数を調べる
	leafNodeLen := minLeafNodeLen(size)
	data := make([]int, leafNodeLen*2)
	// for i := range data {
	// 	// 0<=x<=10^9 なので最大値の比較に影響のない-1を初期値として設定する
	// 	data[i] = -1
	// }

	return &rsq{
		data:        data,
		leafNodeLen: leafNodeLen,
		size:        len(data) - 1,
	}
}

func minLeafNodeLen(size int) int {
	n := 0
	for {
		if pow(2, n) >= size {
			break
		}
		n++
	}
	return pow(2, n)
}

// leafNode内のiの位置をxに更新する
func (rs *rsq) update(i, x int) {
	cellPos := rs.cellPos(i)
	rs.data[cellPos] = x

	currPos := rs.parentIndex(cellPos)
	for currPos >= 1 {
		rs.data[currPos] = max(
			rs.data[rs.leftIndex(currPos)],  // leftMax
			rs.data[rs.rightIndex(currPos)], // rightMax
		)
		currPos = rs.parentIndex(currPos)
	}
}

func (rs *rsq) maxInRange(l, r int) int {
	return rs.maxInRangeHelper(1, rs.leafNodeLen+1, l, r, 1)
}

// [targetL,targetR)区間の最大値を返すメソッド
// [currL,currR)は現在調べている区間
// currCellPosは現在調べているセルのPositionIndex
func (rs *rsq) maxInRangeHelper(currL, currR, targetL, targetR, currCellPos int) int {

	// [currL,currR)が、[targetL,targetR)に含まれるか？の判定によって処理が分かれる

	// 一切含まれない
	if currR <= targetL || currL >= targetR {
		return 0
	}

	// 全部含まれる
	if currL >= targetL && currR <= targetR {
		return rs.data[currCellPos]
	}

	// ↓一部含まれる場合の処理
	mid := (currL + currR) / 2
	leftMax := rs.maxInRangeHelper(currL, mid, targetL, targetR, rs.leftIndex(currCellPos))
	rightMax := rs.maxInRangeHelper(mid, currR, targetL, targetR, rs.rightIndex(currCellPos))

	return max(leftMax, rightMax)
}

func (rs *rsq) leftIndex(i int) int {
	return i * 2
}

func (rs *rsq) rightIndex(i int) int {
	return i*2 + 1
}

func (rs *rsq) parentIndex(i int) int {
	return i / 2
}

func (rs *rsq) cellPos(i int) int {
	return rs.leafNodeLen + i - 1
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	// 長い行を読み込めるようにするための定義
	sc.Buffer([]byte{}, MaxInt)

	N, Q := ni2(sc)

	// セグメント木を構築
	rsq := newRsq(N)

	// クエリを処理
	for range make([]struct{}, Q) {
		q := ni(sc)
		const (
			updateQuery = 1
			answerQuery = 2
		)

		x, y := ni2(sc)

		switch q {
		case updateQuery:
			rsq.update(x, y)
		case answerQuery:
			fmt.Fprintln(writer, rsq.maxInRange(x, y))
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
