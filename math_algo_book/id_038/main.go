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

	N, Q := ni2(sc)

	// A=来場者数のSlice
	A := make([]int, N+1)
	for i := 1; i <= N; i++ {
		A[i]= ni(sc)
	}

	// L=スタート日
	// R=エンド日
	L := make([]int, Q+1)
	R := make([]int, Q+1)
	for i := 1; i <= Q; i++ {
		L[i], R[i] = ni2(sc)
	}

	S := make([]int, N+1)
	S[0] = 0
	S[1] = A[1]

	for i := 2; i < len(S); i++ {
		S[i] = S[i-1]+A[i]
	}

	answer := make([]string, 0, Q)
	for i := 1; i < len(L); i++ {
		start := L[i]
		end := R[i]
		sum := S[end] - S[start-1]
		answer = append(answer, strconv.Itoa(sum))
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
