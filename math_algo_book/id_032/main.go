package main

import (
	"fmt"
	"sort"
)

func main() {
	// n=長さ, x=探す対象
	var n, x int
	fmt.Scan(&n, &x)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

	if contains(a, x) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func contains(list []int, x int) bool {
	var helper func(left, right int) bool
	helper = func(left, right int) bool {
		if left > right {
			return false
		}

		mid := (left + right) / 2
		if list[mid] == x {
			return true
		}

		if x < list[mid] {
			return helper(left, mid-1)
		} else {
			return helper(mid+1, right)
		}
	}

	return helper(0, len(list)-1)
}
