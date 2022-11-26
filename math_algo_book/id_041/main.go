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

	B := make([]int, T+1)

	for i := 0; i < N; i++ {
		l, r := ni2(sc)
		B[l] += 1
		B[r] -= 1
	}

	current := 0
	for i := 0; i < T; i++ {
		current += B[i]
		fmt.Fprintln(writer, current)
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
