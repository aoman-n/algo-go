package main

import (
	"bufio"
	"fmt"
	"io"
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
	T := ni(sc)
	N := ni(sc)

	// 営業時間毎時分の配列を作成->階差を計算する(時間ごと従業員数の階差)
	B := make([]int, T+1)
	for i := 1; i < N+1; i++ {
		l, r := ni2(sc)
		B[l] += 1
		B[r] -= 1
	}

	// A=階差から累積和を計算
	A := make([]int, T)
	A[0] = B[0]
	for i := 1; i < T; i++ {
		A[i] = A[i-1] + B[i]
	}

	for _, v := range A {
		fmt.Fprintln(writer, v)
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
