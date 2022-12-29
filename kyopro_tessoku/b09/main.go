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

const (
	W = 1501
	H = 1501
)

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	N := ni(sc)

	// マス目を作る
	Z := make([][]int, W)
	for i := range Z {
		Z[i] = make([]int, H)
	}

	for range make([]struct{}, N) {
		A, B, C, D := ni4(sc)
		Z[A][B]++
		Z[C][D]++
		Z[C][B]--
		Z[A][D]--
	}

	// 行の累積和
	for x := 1; x < W; x++ {
		for y := 0; y < H; y++ {
			Z[x][y] += Z[x-1][y]
		}
	}

	// 列の累積和
	for x := 0; x < W; x++ {
		for y := 1; y < H; y++ {
			Z[x][y] += Z[x][y-1]
		}
	}

	// 面積の確認
	answer := 0
	for x := 0; x < W; x++ {
		for y := 0; y < H; y++ {
			if Z[x][y] > 0 {
				answer++
			}
		}
	}

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
