package pascal

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// depth=5
//     1        0
//    1 1       1
//   1 2 1      2
//  1 3 3 1     3
// 1 4 6 4 1    4

func GeneratePascalTriangle(depth int32) [][]int {
	result := make([][]int, 0, depth)
	for i := range make([]struct{}, depth) {
		line := make([]int, i+1)
		for ii := range line {
			line[ii] = 1
		}
		result = append(result, line)
	}

	for i := 2; i < len(result); i++ {
		line := result[i]

		for j := 1; j < len(line)-1; j++ {
			beforeLine := result[i-1]
			line[j] = beforeLine[j-1] + beforeLine[j]
		}
	}

	return result
}

// 6文字: **10**
// 7文字: **10***

// len=6, mid=3
// 挿入位置= mid - (対象文字数-1)
// 012345
// ******
// index=2から2文字

// len=7, mid=3
// 挿入位置= mid - (対象文字数-1)
// 0123456
// *******
// index=2から2文字
func center(s string, n int, fill string) string {
	str := ""
	for range make([]struct{}, n) {
		str += fill
	}

	mid := n / 2
	replaceIndex := mid - (len(s) - 1)

	for i := 0; i < len(s); i++ {
		runes := []rune(str)
		target := []rune(s)[i]
		runes[replaceIndex+i] = target
		str = string(runes)
	}

	return str
}

func max(list [][]int) int {
	max := math.MinInt
	for _, line := range list {
		for _, num := range line {
			if num > max {
				max = num
			}
		}
	}

	return max
}

func PrintPascalTriangle(pascalList [][]int) {
	separator := " "
	width := len(strconv.Itoa(max(pascalList))) + (len(pascalList) % 2) + 2

	for i, line := range pascalList {
		lineStr := ""
		beforeSpaceLen := (width / 2) * (len(pascalList) - i)
		lineStr += strings.Repeat(" ", beforeSpaceLen)
		for _, item := range line {
			lineStr += center(strconv.Itoa(item), width, separator)
		}

		fmt.Println(lineStr)
	}
}
