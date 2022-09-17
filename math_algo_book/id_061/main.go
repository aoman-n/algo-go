package main

import "fmt"

const First = "First"
const Second = "Second"

func main() {
	var n int
	fmt.Scan(&n)

	num := 1
	var answer string
	for {
		if num == n {
			answer = Second
			break
		}

		if n <= num {
			answer = First
			break
		}

		num = num*2 + 1
	}

	fmt.Println(answer)
}
