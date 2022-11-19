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

	// N=駅数
	N := ni(sc)
	// A=駅間距離
	A := make([]int, N)
	for i := 1; i < N; i++ {
		A[i] = ni(sc)
	}
	
	// M=行動数
	M := ni(sc)
	// B=行動経路
	B := make([]int, M+1)
	for i := 1; i <= M; i++ {
		B[i] = ni(sc)
	}


	// S=累積和
	S := make([]int, N+1)
	S[0] = 0
	for i := 2; i <= N;i++ {
		S[i] = S[i-1] + A[i-1]
	}

	answer := 0	
	for i := 2; i <= M; i++ {
		l := B[i-1]
		r := B[i]
		if l > r {
			l, r = r, l
		}

		answer += S[r] - S[l]
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
