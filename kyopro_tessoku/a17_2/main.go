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

const MOD = 1000000007

func init() {
	reader = os.Stdin
	writer = os.Stdout
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	N := ni(sc)

	// A=1つ前の部屋からの移動時間
	A := make([]int, N+1)
	// B=2つ前の部屋からの移動時間
	B := make([]int, N+1)

	for i := 2; i <= N; i++ {
		A[i] = ni(sc)
	}
	for i := 3; i <= N; i++ {
		B[i] = ni(sc)
	}

	dp := make([]int, N+1)
	dp[2] = A[2]
	for i := 3; i <= N; i++ {
		dp[i] = min(dp[i-1]+A[i], dp[i-2]+B[i])
	}

	route := make([]int, 0, N+1)
	x := N
	for x > 0 {
		route = append(route, x)
		if dp[x] == dp[x-1]+A[x] {
			x -= 1
		} else {
			x -= 2
		}
	}

	// reverse処理
	for i := len(route)/2 - 1; i >= 0; i-- {
		opp := len(route) - 1 - i
		route[i], route[opp] = route[opp], route[i]
	}

	fmt.Fprintln(writer, len(route))
	for i, v := range route {
		if i == 0 {
			fmt.Fprintf(writer, "%d", v)
		} else {
			fmt.Fprintf(writer, " %d", v)
		}
	}
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
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
