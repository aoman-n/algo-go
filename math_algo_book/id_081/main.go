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
	n := ni(sc)

	answer := 0

	answer += n / 10000
	n = n % 10000

	answer += n / 5000
	n = n % 5000

	answer += n / 1000

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
