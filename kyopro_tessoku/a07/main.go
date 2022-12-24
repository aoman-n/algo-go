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

	// D=開催日数
	// N=参加者数
	D := ni(sc)
	N := ni(sc)

	// 階差
	F := make([]int, D+1)
	for i := 0; i < N; i++ {
		l, r := ni2(sc)
		F[l] += 1
		if (r + 1) < len(F) {
			F[r+1] -= 1
		}
	}

	// 累積和
	S := make([]int, D+1)
	for i := 1; i < D+1; i++ {
		S[i] = S[i-1] + F[i]
	}

	answer := make([]string, 0, D)
	for i := 1; i < D+1; i++ {
		answer = append(answer, strconv.Itoa(S[i]))
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
