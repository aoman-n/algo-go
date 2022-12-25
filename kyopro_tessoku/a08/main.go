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

	H, W := ni2(sc)
	Z := make([][]int, H+1)
	Z[0] = make([]int, W+1)
	for i := 1; i < H+1; i++ {
		row := make([]int, W+1)
		for j := 1; j < W+1; j++ {
			row[j] = ni(sc)
		}
		Z[i] = row
	}

	// 横の累積和を計算
	for rowI := 1; rowI < H+1; rowI++ {
		for colI := 1; colI < W+1; colI++ {
			Z[rowI][colI] = Z[rowI][colI-1] + Z[rowI][colI]
		}
	}

	// 縦の累積和を計算
	for rowI := 1; rowI < H+1; rowI++ {
		for colI := 1; colI < W+1; colI++ {
			Z[rowI][colI] = Z[rowI-1][colI] + Z[rowI][colI]
		}
	}

	Q := ni(sc)
	for range make([]struct{}, Q) {
		A, B, C, D := ni4(sc)
		area := Z[C][D] - Z[A-1][D] - Z[C][B-1] + Z[A-1][B-1]
		fmt.Fprintln(writer, area)
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
