package quiz

import "math"

// https://en.wikipedia.org/wiki/Taxicab_number
// 1729 = Ta(2) = 1**3 + 12**3 = 9**3 + 10**3
// Input 1, 2 => [(1729, [(1, 12), (9, 10)])]
// Input 2, 2 => [(1729, [(1, 12), (9, 10)]), (4104, [(2, 16), (9, 15)])]
// Input 1, 3 => [(87539319, [(167, 436), (228, 423), (255, 414)])]

type cabResult struct {
	answer int
	pair   [][2]int
}

// ex)
// maxAnswerNum = 2
// matchAnswerNum = 1
func taxiCabNumber(maxAnswerNum int, matchAnswerNum int, startAnswer int) []*cabResult {
	gotAnswerCount := 0
	answer := startAnswer
	result := []*cabResult{}

	for gotAnswerCount < maxAnswerNum {
		memo := [][2]int{}
		matchAnswerCount := 0

		maxParam := int(math.Pow(float64(answer), 1.0/3)) + 1

		for x := 1; x < maxParam; x++ {
			for y := x + 1; y < maxParam; y++ {
				if int(math.Pow(float64(x), 3))+int(math.Pow(float64(y), 3)) != answer {
					continue
				}

				memo = append(memo, [2]int{x, y})
				matchAnswerCount++
			}
		}

		if matchAnswerCount == matchAnswerNum {
			gotAnswerCount++
			result = append(result, &cabResult{
				answer: answer,
				pair:   memo,
			})
		}

		answer++
	}

	return result
}
