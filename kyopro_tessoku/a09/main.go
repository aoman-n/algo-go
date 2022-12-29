package main

import (
	"bufio"
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
	H, W, N := ni3(sc)

	// マス目を作る
	Z := make([][]int, W+2)
	for i := range make([]struct{}, len(Z)) {
		Z[i] = make([]int, H+2)
	}

	for range make([]struct{}, N) {
		A, B, C, D := ni4(sc)
		Z[A][B]++
		Z[C+1][D+1]++
		Z[C+1][B]--
		Z[A][D+1]--
	}

	// 行の累積和を計算
	// for x := 1; x <= W+1; x++ {
	// 	for y := 0; y <= H+1; y++ {
	// 		Z[x][y] += Z[x-1][y]
	// 	}
	// }

	// // 列の累積和を計算
	// for x := 0; x <= W+1; x++ {
	// 	for y := 1; y <= H+1; y++ {
	// 		Z[x][y] += Z[x][y-1]
	// 	}
	// }

	// // 出力
	// for y := 1; y <= H; y++ {
	// 	for x := 1; x <= W; x++ {
	// 		cont := ""
	// 		if x >= 2 {
	// 			cont += " "
	// 		}
	// 		cont += strconv.Itoa(Z[x][y])
	// 		fmt.Fprint(writer, cont)
	// 	}
	// 	fmt.Fprintln(writer)
	// }
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
