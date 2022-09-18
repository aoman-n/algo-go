package main

import (
	"bufio"
	"fmt"
	"io"
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
	A, B := ni2(sc)

	answer := exec(A, B)
	fmt.Fprint(writer, answer)
}

func exec(a, b int) int {
	for a != 0 && b != 0 {
		if a >= b {
			a = a % b
		} else {
			b = b % a
		}
	}

	if a == 0 {
		return b
	} else {
		return a
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
