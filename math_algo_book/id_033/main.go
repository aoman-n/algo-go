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

type Pattern int

const (
	Pattern1 Pattern = iota + 1
	Pattern2
	Pattern3
)

func init() {
	reader = os.Stdin
	writer = os.Stdout
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	Ax, Ay := ni2(sc)
	Bx, By := ni2(sc)
	Cx, Cy := ni2(sc)

	// ベクトルの成分を計算
	// BA
	BAx := float64(Ax - Bx)
	BAy := float64(Ay - By)
	// BC
	BCx := float64(Cx - Bx)
	BCy := float64(Cy - By)
	// CA
	CAx := float64(Ax - Cx)
	CAy := float64(Ay - Cy)
	// CB
	CBx := float64(Bx - Cx)
	CBy := float64(By - Cy)

	// パターンを調べる
	// パターン2: 1と3以外
	pattern := Pattern2
	// パターン1: ベクトルBAとベクトルBCの内積が<負>
	if (BAx*BCx)+(BAy*BCy) < 0 {
		pattern = Pattern1
	}
	// パターン3: ベクトルBAとベクトルBCの内積が<正>
	if (CAx*CBx)+(CAy*CBy) < 0 {
		pattern = Pattern3
	}

	var answer float64
	switch pattern {
	case Pattern1:
		answer = math.Sqrt(BAx*BAx + BAy*BAy)
	case Pattern2:
		// ベクトルBAとベクトルBCの外積S(平行四辺形の面積S)
		S := math.Abs(BAx*BCy - BAy*BCx)
		// ベクトルBCの長さ(平行四辺形の底辺の長さBCLength)
		BCLength := math.Sqrt(BCx*BCx + BCy*BCy)
		// (平行四辺形の面積S) / (平行四辺形の底辺の長さBCLength) = (平行四辺形の高さ = ベクトルAHの長さ)
		answer = S / BCLength
	case Pattern3:
		answer = math.Sqrt(CAx*CAx + CAy*CAy)
	}

	// fmt.Fprint(writer, answer)
	fmt.Fprintf(writer, "%.12f", answer)
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
