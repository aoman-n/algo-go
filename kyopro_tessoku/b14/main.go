package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
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
	N, K := ni2(sc)
	A := nis(sc, N)

	mid := len(A) / 2
	P := getAllSumBox(A[:mid])
	Q := getAllSumBox(A[mid:])

	// PorQにKが存在するか確認
	for _, v := range P {
		if v == K {
			fmt.Fprint(writer, "Yes")
			return
		}
	}
	for _, v := range Q {
		if v == K {
			fmt.Fprint(writer, "Yes")
			return
		}
	}

	sort.Ints(Q)

	for _, p := range P {
		if search(Q, K-p) {
			fmt.Fprint(writer, "Yes")
			return
		}
	}

	fmt.Fprint(writer, "No")
}

// 全探索ですべてのケースの合計値をスライスに格納して返す
func getAllSumBox(a []int) []int {
	allCount := pow(2, len(a))
	sumBox := make([]int, 0, allCount)

	for i := 0; i < allCount; i++ {
		sum := 0
		for j := 0; j < len(a); j++ {
			if (pow(2, j) & i) != 0 {
				sum += a[j]
			}
		}
		sumBox = append(sumBox, sum)
	}

	return sumBox
}

func search(a []int, target int) bool {
	l, r := 0, len(a)-1

	for r >= l {
		mid := (l + r) / 2

		if a[mid] == target {
			return true
		}

		if target > a[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
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
