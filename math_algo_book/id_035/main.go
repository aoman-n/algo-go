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

// x1,y1,r1
// x2,y2,r2

func vLength(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	x1, y1, r1 := ni3(sc)
	x2, y2, r2 := ni3(sc)

	// (x1, y1) -> (x2, y2)
	Vx := x2 - x1
	Vy := y2 - y1
	// v=2つの円の中心から中心の距離
	v := math.Abs(vLength(float64(Vx), float64(Vy)))

	var smallR float64
	var largeR float64
	if r1 > r2 {
		smallR = float64(r2)
		largeR = float64(r1)
	} else {
		smallR = float64(r1)
		largeR = float64(r2)
	}

	answer := 1
	if smallR+largeR < v {
		answer = 5
	} else if smallR+largeR == v {
		answer = 4
	} else {
		if smallR+v > largeR {
			answer = 3
		} else if smallR+v == largeR {
			answer = 2
		} else {
			answer = 1
		}
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

func ni3(sc *bufio.Scanner) (int, int, int) {
	return ni(sc), ni(sc), ni(sc)
}
