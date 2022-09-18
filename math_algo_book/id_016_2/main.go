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

func init() {
	reader = os.Stdin
	writer = os.Stdout
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	N := ni(sc)
	A := nis(sc, N)

	for i := 0; i < len(A)-1; i++ {
		A[i+1] = greatestCommonDivisor(A[i], A[i+1])
	}

	answer := A[len(A)-1]
	fmt.Fprint(writer, answer)
}

func greatestCommonDivisor(a, b int) int {
	for a != 0 && b != 0 {
		if a >= b {
			a = a % b
		} else {
			b = b % a
		}
	}

	if a == 0 {
		return b
	} else {
		return a
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

func nis(sc *bufio.Scanner, n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni(sc)
	}
	return a
}
