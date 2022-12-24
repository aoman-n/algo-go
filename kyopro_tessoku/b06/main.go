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
	// N=くじを引いた回数
	N := ni(sc)
	// A=n回目のくじのあたりはずれ配列
	A := make([]int, N+1)
	for i := 1; i < N+1; i++ {
		A[i] = ni(sc)
	}

	// hitS=当たりの累積和
	// lostS=当たりの累積和
	hitS := make([]int, N+1)
	lostS := make([]int, N+1)
	// 1が当たり
	hitNumber := 1
	if A[1] == hitNumber {
		hitS[1] = 1
	} else {
		lostS[1] = 1
	}
	for i := 2; i < N+1; i++ {
		if A[i] == hitNumber {
			hitS[i] = hitS[i-1] + 1
			lostS[i] = lostS[i-1]
		} else {
			hitS[i] = hitS[i-1]
			lostS[i] = lostS[i-1] + 1
		}
	}

	Q := ni(sc)
	answer := make([]string, Q)
	for i := 0; i < Q; i++ {
		l, r := ni2(sc)
		hitNum := hitS[r] - hitS[l-1]
		lostNum := lostS[r] - lostS[l-1]
		switch {
		case hitNum > lostNum:
			answer[i] = "win"
		case hitNum < lostNum:
			answer[i] = "lose"
		default:
			answer[i] = "draw"
		}
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
