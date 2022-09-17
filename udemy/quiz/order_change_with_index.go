package quiz

import "strings"

/*
Input : ['h', 'y', 'n', 'p', 't', 'o'], [3, 1, 4, 0, 5, 2]
Output: python
*/

// 新しくsliceを作る方法
func orderChangeByIndexesV1(chars []string, indexes []int) string {
	resultChars := make([]string, len(indexes))

	for i, v := range indexes {
		resultChars[i] = chars[v]
	}

	return strings.Join(resultChars, "")
}

// 新しくsliceを作ってメモリを確保しない方法
// ['h', 'y', 'n', 'p', 't', 'o']
// [  3,   1,   4,   0,   5,   2]
// =  p    y    t    h    o    n

// i=0
// ['h', 'y', 'n', 'p', 't', 'o']
// [  3,   1,   4,   0,   5,   2]
// ['p', 'y', 'n', 'h', 't', 'o']
// [  0,   1,   4,   3,   5,   2]

// i=1
// ['p', 'y', 'n', 'h', 't', 'o']
// [  0,   1,   4,   3,   5,   2]

// i=2
// ['p', 'y', 't', 'h', 'n', 'o']
// [  0,   1,   5,   3,  4,   2]

// [  0,   1,   4,   3,   5,   2] v=4
// [  0,   1,   5,   3,   4,   2] v=5
// [  0,   1,   2,   3,   4,   5] v=2

// i=3
// [  0,   1,   2,   3,   4,   5]

// i=4
// [  0,   1,   2,   3,   4,   5]

// i=4
// [  0,   1,   2,   3,   4,   5]
func orderChangeByIndexesV2(chars []string, indexes []int) string {
	i, lenIndexes := 0, len(indexes)-1
	for i < lenIndexes {
		for indexes[i] != i {
			v := indexes[i]
			indexes[i], indexes[v] = indexes[v], indexes[i]
			chars[i], chars[v] = chars[v], chars[i]
		}
		i++

	}

	return strings.Join(chars, "")
}
