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

type position struct {
	x, y int
}

func newPosition(x, y int) *position {
	return &position{
		x: x,
		y: y,
	}
}

type queue struct {
	list *list.List
}

func newQueue() *queue {
	return &queue{
		list: list.New(),
	}
}

func (q *queue) enqueue(pos *position) {
	q.list.PushBack(pos)
}

func (q *queue) dequeue() *position {
	el := q.list.Front()
	if el == nil {
		return nil
	}

	q.list.Remove(el)
	r := el.Value.(*position)
	return r
}

func (q *queue) len() int {
	return q.list.Len()
}

const (
	initialCount = -1
)

func getInitialRowCount(size int) []int {
	ret := make([]int, size)
	for i := range ret {
		ret[i] = initialCount
	}
	return ret
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	// 長い行を読み込めるようにするための定義
	sc.Buffer([]byte{}, MaxInt)

	rowCount, columnCount := ni2(sc)
	startY, startX := ni2(sc)
	goalY, goalX := ni2(sc)

	// 入力例)
	// ########
	// #......#
	// #.######
	// #..#...#
	// #..##..#
	// ##.....#
	// ########

	// maze: 迷路情報を保持する変数
	// row,column両方とも1Index
	// 0Indexは壁として利用
	maze := make([][]string, rowCount+1)
	for i := 0; i < len(maze); i++ {
		if i == 0 {
			maze[i] = make([]string, columnCount+1)
		} else {
			s := ns(sc)
			maze[i] = split(s)
		}
	}

	// counts: 各マスへの移動数を保持する変数
	counts := make([][]int, rowCount+1)
	for i := range counts {
		counts[i] = getInitialRowCount(columnCount + 1)
	}

	q := newQueue()
	q.enqueue(newPosition(startX, startY))
	counts[startY][startX] = 0

	for q.len() > 0 {
		currPos := q.dequeue()
		nextPositions := nextPositions(currPos, rowCount, columnCount)

		// 次のマスのカウント更新 + エンキュー処理
		currCount := counts[currPos.y][currPos.x]
		nextCount := currCount + 1
		for _, p := range nextPositions {
			// 進めるポジションかの判定
			x, y := p.x, p.y
			if maze[y][x] == "#" {
				continue
			}
			if counts[y][x] != initialCount {
				continue
			}

			// 更新処理
			counts[p.y][p.x] = nextCount
			q.enqueue(p)
		}

		if counts[goalY][goalX] != initialCount {
			break
		}
	}

	fmt.Fprintln(writer, counts[goalY][goalX])
}

func nextPositions(currPos *position, rowCount, columnCount int) []*position {
	nextPositions := []*position{
		newPosition(currPos.x, currPos.y-1), // up
		newPosition(currPos.x, currPos.y+1), // down
		newPosition(currPos.x-1, currPos.y), // left
		newPosition(currPos.x+1, currPos.y), // right
	}

	ret := make([]*position, 0, 4)
	for _, p := range nextPositions {
		if inRange(p, rowCount, columnCount) {
			ret = append(ret, p)
		}
	}

	return ret
}

func inRange(pos *position, rowCount, columnCount int) bool {
	if pos.x <= 0 || pos.y <= 0 {
		return false
	}
	if pos.y > rowCount || pos.x > columnCount {
		return false
	}

	return true
}

func split(s string) []string {
	ret := make([]string, len(s)+1)
	for i, v := range s {
		ret[i+1] = string(v)
	}
	return ret
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
