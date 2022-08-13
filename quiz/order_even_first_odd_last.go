// l = [1, 3, 4, 2, 4, 5, 1, 6, 9, 8]  =>  l = [4, 2, 4, 6, 8, 1, 3, 5, 1, 9]
package quiz

func orderEvenFirstOddLastV1(numbers []int) []int {
	even, odd := []int{}, []int{}

	for _, n := range numbers {
		if n%2 == 0 {
			even = append(even, n)
		} else {
			odd = append(odd, n)
		}
	}

	return append(even, odd...)
}

func orderEvenFirstOddLastV2(numbers []int) {
	start, end := 0, len(numbers)-1

	for start < end {
		if numbers[start]%2 == 0 {
			start++
		} else {
			numbers[start], numbers[end] = numbers[end], numbers[start]
			end--
		}
	}
}
