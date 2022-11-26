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

	// 例) N=10
	N := ni(sc)

	// N個分の容器を作成
	S := make([]bool, N+1)
	for i := range S {
		S[i] = true
	}

	for i := 2; i * i <= N; i++ {
		if !S[i] {
			continue
		}

		for j := i * 2; j <= N; j += i {
			S[j] = false
		}
	}

	// 素数の配列を作成
	ret := make([]string, 0, N)
	for i := 2; i < len(S); i++ {
		if S[i] {
			ret = append(ret, strconv.Itoa(i))
		}
	}

	fmt.Fprint(writer, strings.Join(ret, " "))
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
