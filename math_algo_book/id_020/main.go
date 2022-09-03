package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	numbers := make([]int, n)
	for i := range numbers {
		fmt.Scan(&numbers[i])
	}

	// 5枚のカードを選んで、和が1000になる通り
	// i, j, k, l, m

	var answer int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				for l := k + 1; l < n; l++ {
					for m := l + 1; m < n; m++ {
						if numbers[i]+numbers[j]+numbers[k]+numbers[l]+numbers[m] == 1000 {
							answer++
						}
					}
				}
			}
		}
	}

	fmt.Println(answer)
}
