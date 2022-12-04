package quiz

import (
	"fmt"

	"github.com/aoman-n/algo-go/udemy/pkg/smath"
)

// {1, -2, 3, 6, -1, 2, 4, -5, 2}
// 14 {3, 6, -1, 2, 4}

func GetMaxSequenceSumBad(numbers []int) int {
	results := []int{}

	for i, v := range numbers {
		sums := []int{v}
		for j := i + 1; j < len(numbers); j++ {
			sum := sums[len(sums)-1] + numbers[j]
			sums = append(sums, sum)
		}
		fmt.Printf("sums: %v\n", sums)
		results = append(results, sums...)
	}

	return smath.MaxInNumbers(results)
}

func GetMaxSequenceSum(numbers []int, operator func(int, int) int) int {
	resultSequence, sumSequence := 0, 0
	for _, num := range numbers {
		sumSequence = operator(num, sumSequence+num)
		resultSequence = operator(resultSequence, sumSequence)
	}

	return resultSequence
}

func FindMaxCircularSequenceSum(numbers []int) int {
	maxSequenceSum := GetMaxSequenceSum(numbers, smath.MaxInt)
	maxWrapSequenceSum := smath.SumInts(numbers) - GetMaxSequenceSum(numbers, smath.MinInt)

	return smath.MaxInt(maxSequenceSum, maxWrapSequenceSum)
}
