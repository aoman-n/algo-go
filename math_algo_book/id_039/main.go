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

	// N=区画数
	// Q=積雪日数
	N, Q := ni2(sc)

	// L=開始区画,R=終了区画,X=積雪量
	L := make([]int, Q+1)
	R := make([]int, Q+1)
	X := make([]int, Q+1)

	for i := 1; i <= Q; i++ {
		L[i], R[i], X[i] = ni3(sc)
	}

	// B=階差を格納
	// 計算しやすいように下記にする
	// index=0とN+1には使用しない値を入れておく
	B := make([]int, N+1)

	for i := 1; i <= Q; i++ {
		l, r, x := L[i], R[i], X[i]
		B[l] += x
		// ここで+1するので、Bの配列を1伸ばしている
		if r+1 < len(B) {
			B[r+1] -= x
		}
	}

	answer := ""

	for i := 2; i <= N; i++ {
		if B[i] > 0 {
			answer += "<"
		} else if B[i] == 0 {
			answer += "="
		} else {
			answer += ">"
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
