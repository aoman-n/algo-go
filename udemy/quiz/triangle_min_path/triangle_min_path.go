package triangle

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// input: 5, 20
//     1        0
//    1 1       1
//   1 2 1      2
//  1 3 3 1     3
// 1 4 6 4 1    4

func GenerateRandTriangle(depth, maxNum int) [][]int {
	result := make([][]int, 0, depth)

	for i := 1; i <= depth; i++ {
		line := make([]int, i)
		for j := range line {
			line[j] = rand.Intn(maxNum)
		}
		result = append(result, line)
	}

	return result
}

func max(list [][]int) int {
	max := math.MinInt
	for _, line := range list {
		for _, val := range line {
			if val > max {
				max = val
			}
		}
	}

	return max
}

// n=4
// mid=2
// numLen=2
// startIndex=min-(numLen-1)
//
// 0123
// ****

// n=5
// mid=2
// numLen=2
// startIndex=mid-(numLen-1)
//
// 01234
// *****

func center(s string, n int, fill string) string {
	str := ""
	for range make([]struct{}, n) {
		str += fill
	}

	mid := n / 2
	sLen := len(s)
	startIndex := mid - (sLen - 1)

	for i := 0; i < sLen; i++ {
		replaceIndex := startIndex + i
		runes := []rune(str)
		runes[replaceIndex] = []rune(s)[i]
		str = string(runes)
	}

	return str
}

func Print(list [][]int) {
	separator := " "
	width := len(strconv.Itoa(max(list))) + (len(list) % 2) + 2

	for i, line := range list {
		leftSpace := (width / 2) * (len(list) - i)
		str := strings.Repeat(" ", leftSpace)
		for _, val := range line {
			str += center(strconv.Itoa(val), width, separator)
		}
		fmt.Println(str)
	}
}

func SumMinPath(list [][]int) int {
	treeSum := make([][]int, len(list))
	copy(treeSum, list)

	for i := 1; i < len(treeSum); i++ {
		currentLine := treeSum[i]
		for j, val := range currentLine {
			var sum int
			switch {
			case j == 0:
				sum = val + treeSum[i-1][0]
			case j == len(currentLine)-1:
				sum = val + treeSum[i-1][len(treeSum[i-1])-1]
			default:
				left := treeSum[i-1][j-1]
				right := treeSum[i-1][j]
				smaller := int(math.Min(float64(left), float64(right)))
				sum = val + smaller
			}
			currentLine[j] = sum
		}
	}

	Print(treeSum)

	lastLine := treeSum[len(treeSum)-1]

	min := math.MaxInt
	for _, val := range lastLine {
		if val < min {
			min = val
		}
	}

	return min
}
