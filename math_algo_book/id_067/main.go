package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var reader io.Reader
var writer io.Writer

func init() {
	reader = os.Stdin
	writer = os.Stdout
}

// squares=
// [
//     0  1  2  3
// 	0 [1, 1, 1, 1],
//	1 [1, 1, 1, 1],
// 	2 [1, 1, 1, 1],
//  3 [1, 1, 1, 1],
// ]

// columnSums [4, 4, 4, 4]
// rowSums    [4, 4, 4, 4]

// row-column
// 0-0         = rowSums[0] + columnSums[0] - squares[0][0]
// 0-1         = rowSums[0] + columnSums[1] - squares[0][1]
// 0-2				 = rowSums[0] + columnSums[2] - squares[0][2]
// 0-3				 = rowSums[0] + columnSums[3] - squares[0][3]
// 1-0				 = rowSums[1] + columnSums[0] - squares[1][0]
// 1-1				 = rowSums[1] + columnSums[1] - squares[1][1]
// ...

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	H, W := ni2(sc)
	A := make([][]int, H)
	for i := range A {
		A[i] = nis(sc, W)
	}

	columnSums := make([]int, W)
	for i := range columnSums {
		sum := 0
		for _, row := range A {
			sum += row[i]
		}
		columnSums[i] = sum
	}

	rowSums := make([]int, H)
	for i := range rowSums {
		row := A[i]
		sum := 0
		for _, v := range row {
			sum += v
		}
		rowSums[i] = sum
	}

	sumsSquare := make([][]int, H)
	for rowIndex := range sumsSquare {
		sumLine := make([]int, W)
		for columnIndex := range sumLine {
			sumLine[columnIndex] = columnSums[columnIndex] + rowSums[rowIndex] - A[rowIndex][columnIndex]
		}
		sumsSquare[rowIndex] = sumLine
	}

	for _, line := range sumsSquare {
		fmt.Fprint(writer, strings.Join(toStrings((line)), " ")+"\n")
	}
}

func toStrings(a []int) []string {
	ret := make([]string, len(a))
	for i, n := range a {
		ret[i] = strconv.Itoa(n)
	}
	return ret
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

func nis(sc *bufio.Scanner, n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni(sc)
	}
	return a
}
