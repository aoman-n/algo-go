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
	// N=部屋数
	N := ni(sc)
	// A=部屋ごとの最大収容人数
	A := make([]int, N+1)
	for i := 1; i < len(A); i++ {
		A[i] = ni(sc)
	}

	lMax := make([]int, N+2)
	for i := 1; i <= N; i++ {
		lMax[i] = max(lMax[i-1], A[i])
	}

	rMax := make([]int, N+2)
	for i := N; i > 0; i-- {
		rMax[i] = max(rMax[i+1], A[i])
	}

	// D=工事日数
	D := ni(sc)
	for range make([]struct{}, D) {
		l, r := ni2(sc)
		fmt.Fprintln(writer, max(lMax[l-1], rMax[r+1]))
	}
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
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
