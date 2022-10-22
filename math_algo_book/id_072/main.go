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

func hasOverTwo(t, a, b int) bool {
	// a~b範囲の"最大の倍数"を計算(小数点切り捨てることで計算する)
	left := b / t * t
	// a~b範囲の"最小の倍数"を計算(小数点を切り上げることで計算する)
	right := int(math.Ceil(float64(a)/float64(t)) * float64(t))

	return (left - right) >= 1
}

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	// A<=x,y<=B
	A, B := ni2(sc)

	answer := 0
	for i := 1; i <= B; i++ {
		if hasOverTwo(i, A, B) {
			answer = i
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

func ni2(sc *bufio.Scanner) (int, int) {
	return ni(sc), ni(sc)
}
