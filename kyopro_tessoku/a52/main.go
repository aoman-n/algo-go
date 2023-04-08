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

// <queries>
// 1. 最後尾に並ぶ
// 2. 先頭を答える
// 3. 先頭を抜ける

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	// 長い行を読み込めるようにするための定義
	sc.Buffer([]byte{}, MaxInt)

	N := ni(sc)

	q := newStringQueue()

	for range make([]struct{}, N) {
		query := ni(sc)

		switch query {
		case 1:
			user := ns(sc)
			q.enqueue(user)
		case 2:
			fmt.Fprintln(writer, q.peekFront())
		case 3:
			q.dequeue()
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

type node struct {
	val  interface{}
	next *node
}

func newNode(val interface{}) *node {
	return &node{
		val:  val,
		next: nil,
	}
}

type stringQueue struct {
	back  *node
	front *node
}

func newStringQueue() *stringQueue {
	return &stringQueue{
		back:  nil,
		front: nil,
	}
}

func (q *stringQueue) enqueue(str string) {
	newNode := newNode(str)

	if q.front == nil {
		q.front = newNode
		q.back = newNode
	} else {
		q.back.next = newNode
		q.back = newNode
	}
}

func (q *stringQueue) dequeue() string {
	if q.front == nil {
		return ""
	}

	tmpNode := q.front
	q.front = q.front.next
	return tmpNode.val.(string)
}

func (q *stringQueue) peekFront() string {
	if q.front == nil {
		return ""
	}

	return q.front.val.(string)
}

func (q *stringQueue) peekBack() string {
	if q.back == nil {
		return ""
	}

	return q.back.val.(string)
}
