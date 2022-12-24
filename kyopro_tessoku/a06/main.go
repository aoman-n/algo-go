package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
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

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	// N=日数
	// Q=質問の数
	N, Q := ni2(sc)

	// 0は捨てメモリ
	// 1日目から取得する
	A := make([]int, N+1)
	for i := 1; i < N+1; i++ {
		A[i] = ni(sc)
	}

	// S=累積和
	S := make([]int, N+1)
	S[1] = A[1]
	for i := 2; i < N+1; i++ {
		S[i] = S[i-1] + A[i]
	}

	// Q回ループして答えを出力していく
	answer := make([]string, Q)
	for i := 0; i < Q; i++ {
		l, r := ni2(sc)
		answer[i] = strconv.Itoa(S[r] - S[l-1])
	}

	fmt.Fprint(writer, strings.Join(answer, "\n"))
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
