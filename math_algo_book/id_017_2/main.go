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

// 最小公倍数 = a*b / aとbの最大公約数
func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	N := ni(sc)
	A := nis(sc, N)

	answer := leastCommonMultiple(A[0], A[1])
	for i := 2; i < len(A); i++ {
		answer = leastCommonMultiple(answer, A[i])
	}
	fmt.Fprint(writer, answer)
}

func gcd(a, b int) int {
	for a >= 1 && b >= 1 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}

	if a >= 1 {
		return a
	}
	return b
}

func leastCommonMultiple(a, b int) int {
	return (a / gcd(a, b)) * b
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
