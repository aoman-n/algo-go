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

	// N=点の数
	// X,Y=座標の配列
	N := ni(sc)
	X := make([]int, N)
	Y := make([]int, N)
	for i := 0; i < N; i++ { 
		X[i], Y[i] = ni2(sc)
	}

	// 累積和を入れる表を初期化
	Z := make([][]int, 1500+1)
	for i := range Z {
		Z[i] = make([]int, 1500+1)
	}

	// 各点を調べZに入れていく
	for i := 0; i < N; i++ {
		Z[Y[i]][X[i]]++
	}

	// 横の累積和を計算
	for rowI := 0; rowI < 1500+1; rowI++ {
		for colI := 1; colI < 1500+1; colI++ {
			Z[rowI][colI] = Z[rowI][colI-1] + Z[rowI][colI]
		}
	}

	// 縦の累積和を計算
	for colI := 0; colI < 1500+1; colI++ {
		for rowI := 1; rowI < 1500+1; rowI++ {
			Z[rowI][colI] = Z[rowI-1][colI] + Z[rowI][colI]
		}
	}

	Q := ni(sc)
	for range make([]struct{}, Q) {
		a, b, c, d := ni4(sc)
		sum := Z[d][c]
		switch {
		case a <= 0 || b <= 0:
			fmt.Fprintln(writer, sum)
		default:
			answer := sum - Z[d][a-1] - Z[b-1][c] + Z[b-1][a-1]
			fmt.Fprintln(writer, answer)
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
