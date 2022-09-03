package main

import "fmt"

type Color int

const (
	ColorRed    Color = 1
	ColorYellow Color = 2
	ColorBlue   Color = 3
)

func main() {
	var n int
	fmt.Scan(&n)

	colorCounts := map[Color]int{
		ColorRed:    0,
		ColorYellow: 0,
		ColorBlue:   0,
	}

	for range make([]int, n) {
		var a int
		fmt.Scan(&a)
		colorCounts[Color(a)]++
	}

	var answer int
	for _, count := range colorCounts {
		answer += (count * (count - 1)) / 2
	}

	fmt.Println(answer)
}
