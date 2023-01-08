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
	N, K := ni2(sc)

	A := make([]int, N+1)
	for i := 1; i < len(A); i++ {
		A[i] = ni(sc)
	}

	sum := sumFactory(A)

	// しゃくとり法
	R := make([]int, len(A))
	for i := 1; i < len(A); i++ {
		if i != 1 {
			R[i] = R[i-1]
		}

		for R[i] < N && sum(i, R[i]+1) <= K {
			R[i]++
		}
	}

	answer := 0
	for i := 1; i < len(R); i++ {
		answer += R[i] - i + 1
	}

	fmt.Fprint(writer, answer)
}

func sumFactory(A []int) func(l, r int) int {
	// 累積和を求める
	S := make([]int, len(A))
	for i := 1; i < len(S); i++ {
		S[i] = S[i-1] + A[i]
	}

	return func(l, r int) int {
		return S[r] - S[l-1]
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
