package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
	N, K := ni2(sc)
	A := nis(sc, K)

	answer := 0
	// ビット全探索
	for i := 1; i < int(math.Pow(2, float64(K))); i++ {
		count := 0
		lcmNumber := 1

		// j=桁目
		for j := 1; j <= K; j++ {
			if i&int(math.Pow(2, float64(j-1))) != 0 {
				count += 1
				lcmNumber = lcm(lcmNumber, A[j-1])
			}
		}

		if count%2 == 0 {
			answer -= N / lcmNumber
		} else {
			answer += N / lcmNumber
		}
	}

	fmt.Fprint(writer, answer)
}

func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
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
