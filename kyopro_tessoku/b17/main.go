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
	H := make([]int, N+1)
	for i := 1; i < N+1; i++ {
		H[i] = ni(sc)
	}

	dp := make([]int, N+1)
	dp[2] = abs(H[1] - H[2])
	for i := 3; i < N+1; i++ {
		beforeOne := dp[i-1] + abs(H[i]-H[i-1])
		beforeTwo := dp[i-2] + abs(H[i]-H[i-2])
		dp[i] = min(beforeOne, beforeTwo)
	}

	route := make([]int, 0, N)
	x := N
	for {
		route = append(route, x)

		if x <= 1 {
			break
		}

		// 2つ前から判定することで、より少ない回数の移動ルートを探す
		if dp[x] == dp[x-2]+abs(H[x]-H[x-2]) {
			x -= 2
		} else {
			x -= 1
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

func abs(n int) int {
	return int(math.Abs(float64(n)))
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
