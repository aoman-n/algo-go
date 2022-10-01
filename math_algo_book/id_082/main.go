package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	n := ni(sc)
	a := nir(sc, n, 2)

	sort.Slice(a, func(i, j int) bool {
		if a[i][1] == a[j][1] {
			return a[i][0] < a[j][0]
		} else {
			return a[i][1] < a[j][1]
		}
	})

	answer := 0
	currentTime := 0
	for i := 0; i < len(a); i++ {
		if a[i][0] < currentTime {
			continue
		}
		answer++              // 映画をカウント
		currentTime = a[i][1] // 見た映画の終了時刻を入れる
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
