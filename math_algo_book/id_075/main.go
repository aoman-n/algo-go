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

func f(n int) int {
	ret := n
	for i := n - 1; i >= 1; i-- {
		ret *= i
	}
	return ret
}

func ncr(r, n int) int {
	if n == 0 || r-n == 0 {
		return 1
	}

	return (f(r) / (f(n) * f(r-n))) % MOD
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	N := ni(sc)
	A := make([]int, N+1)
	for i := 1; i <= N; i++ {
		A[i] = ni(sc)
	}

	answer := 0
	for i := 1; i <= N; i++ {
		answer += A[i] * ncr(N-1, i-1)
		answer %= MOD
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
