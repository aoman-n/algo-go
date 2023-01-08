package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var reader io.Reader
var writer io.Writer

const MOD = 1000000007

func init() {
	reader = os.Stdin
	writer = os.Stdout
}

type positions struct {
	val map[int][]int
}

func (p *positions) getIndexes(v int) []int {
	return p.val[v]
}

// result = map[targetInt][]index
func newPositions(a []int) *positions {
	m := make(map[int][]int)

	for i, v := range a {
		m[v] = append(m[v], i)
	}

	return &positions{m}
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	N := ni(sc)
	A := nis(sc, N)

	positions := newPositions(A)

	Q := uniq(A)
	sort.Ints(Q)

	answer := make([]string, len(A))

	for i, target := range Q {
		val := i + 1

		indexes := positions.getIndexes(target)
		for _, index := range indexes {
			answer[index] = strconv.Itoa(val)
		}
	}

	fmt.Fprint(writer, strings.Join(answer, " "))
}

func uniq(a []int) []int {
	m := make(map[int]struct{})

	for _, v := range a {
		m[v] = struct{}{}
	}

	ret := make([]int, 0, len(m))
	for key := range m {
		ret = append(ret, key)
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
