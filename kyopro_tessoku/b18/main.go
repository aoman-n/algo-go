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
	N, S := ni2(sc)
	A := make([]int, N+1)
	for i := 1; i < N+1; i++ {
		A[i] = ni(sc)
	}

	// [縦軸][横軸]int
	dp := make([][]bool, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]bool, S+1)
	}

	dp[0][0] = true

	// i=縦軸=カード番号
	for i := 1; i < N+1; i++ {
		// j=横軸=合計数値
		for j := 0; j < S+1; j++ {
			if j < A[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				if dp[i-1][j] || dp[i-1][j-A[i]] {
					dp[i][j] = true
				}
			}
		}
	}

	// 合計値がSとなるパターンが存在しない場合は-1を出力
	if !dp[N][S] {
		fmt.Fprint(writer, "-1")
		return
	}

	selected := make([]int, 0, N)
	j := S
	for i := N; i >= 1; i-- {
		if A[i] > j {
			// 引けないのでそのまま上↑にスライド
		} else {
			if dp[i-1][j-A[i]] {
				j = j - A[i]
				selected = append(selected, i)
			}
		}
	}

	// reverse処理
	for i := len(selected)/2 - 1; i >= 0; i-- {
		opp := len(selected) - 1 - i
		selected[i], selected[opp] = selected[opp], selected[i]
	}

	fmt.Fprintln(writer, len(selected))

	for i, v := range selected {
		if i == 0 {
			fmt.Fprintf(writer, "%d", v)
		} else {
			fmt.Fprintf(writer, " %d", v)
		}
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
