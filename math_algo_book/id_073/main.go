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

const mod = 1000000007

func main() {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	n := ni(sc)
	a := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		a[i] = ni(sc)
	}

	answer := 0
	for i := n; i >= 1; i-- {
		answer += (int(math.Pow(2, float64(i-1))) * a[i]) % mod
	}

	fmt.Fprint(writer, answer%mod)
}

// 計算量O(n**2-1)のパターン
// func main() {
// 	sc := bufio.NewScanner(reader)
// 	sc.Split(bufio.ScanWords)
// 	n := ni(sc)
// 	a := make([]int, n+1)
// 	for i := 1; i < n+1; i++ {
// 		a[i] = ni(sc)
// 	}

// 	answer := [][]int{}
// 	for i := 1; i < int(math.Pow(2, float64(n))); i++ {
// 		selectedNumbers := []int{}
// 		for j := 1; j <= n; j++ {
// 			if i&int(math.Pow(2, float64(j-1))) != 0 {
// 				selectedNumbers = append(selectedNumbers, a[j])
// 			}
// 		}
// 		answer = append(answer, selectedNumbers)
// 	}

// 	fmt.Printf("answer: %v\n", answer)

// 	sum := 0
// 	for _, v := range answer {
// 		sum += max(v...)
// 	}

// 	fmt.Printf("sum: %v\n", sum)
// }

// func max(a ...int) int {
// 	maxNumber := a[0]
// 	for i := 1; i < len(a); i++ {
// 		if a[i] > maxNumber {
// 			maxNumber = a[i]
// 		}
// 	}
// 	return maxNumber
// }

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
