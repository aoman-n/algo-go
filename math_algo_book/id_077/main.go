package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

var reader io.Reader
var writer io.Writer

func init() {
	reader = os.Stdin
	writer = os.Stdout
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	N := ni(sc)
	A := nir(sc, N, 2)

	X := make([]int, N)
	Y := make([]int, N)
	for i, c := range A {
		X[i] = c[0]
		Y[i] = c[1]
	}

	sort.Ints(X)
	sort.Ints(Y)

	sumX := 0
	sumY := 0
	for i := 0; i < N; i++ {
		curr := i + 1
		sumX += X[i] * (N - 2*curr + 1)
		sumY += Y[i] * (N - 2*curr + 1)
	}

	answer := abs(sumX) + abs(sumY)
	fmt.Fprint(writer, answer)
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
