package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	input  io.Reader
	output io.Writer
)

func init() {
	input = os.Stdin
	output = os.Stdout
}

func main() {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		numbers[i], _ = strconv.Atoi(scanner.Text())
	}

	mergeSort(numbers)

	answer := make([]string, len(numbers))
	for i, val := range numbers {
		answer[i] = strconv.Itoa(val)
	}

	fmt.Fprintln(output, strings.Join(answer, " "))
}

func mergeSort(numbers []int) {
	if len(numbers) <= 1 {
		return
	}

	mid := len(numbers) / 2

	left := make([]int, len(numbers[:mid]))
	copy(left, numbers[:mid])
	right := make([]int, len(numbers[mid:]))
	copy(right, numbers[mid:])

	mergeSort(left)
	mergeSort(right)

	lIndex, rIndex, retIndex := 0, 0, 0

	for lIndex < len(left) && rIndex < len(right) {
		if left[lIndex] <= right[rIndex] {
			numbers[retIndex] = left[lIndex]
			lIndex++
		} else {
			numbers[retIndex] = right[rIndex]
			rIndex++
		}
		retIndex++
	}

	for lIndex < len(left) {
		numbers[retIndex] = left[lIndex]
		lIndex++
		retIndex++
	}

	for rIndex < len(right) {
		numbers[retIndex] = right[rIndex]
		rIndex++
		retIndex++
	}
}
