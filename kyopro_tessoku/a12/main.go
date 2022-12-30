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

const MOD = 1000000007

func init() {
	reader = os.Stdin
	writer = os.Stdout
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	// N=プリンターの台数
	// K=目標プリント枚数
	// A=何秒毎にプリントできるか
	N, K := ni2(sc)
	A := make([]int, N)
	min := int(math.Pow(10, 9) + 1)
	for i := 0; i < len(A); i++ {
		A[i] = ni(sc)
		if A[i] < min {
			min = A[i]
		}
	}

	r := K * min
	l := 0
	for l < r {
		mid := (l + r) / 2
		sum := sumSheets(mid, A)
		if sum >= K {
			r = mid // mid-1ではなくmidを入れる必要がある
		} else {
			l = mid + 1
		}
	}

	fmt.Fprint(writer, l)
}

func sumSheets(sec int, printers []int) int {
	sum := 0
	for _, v := range printers {
		sum += sec / v
	}
	return sum
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
