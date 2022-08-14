package quiz

import "math"

// -----------------------------
// 素数抜き出し問題
// Input: 50 =>	 Output: [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47]

// 冗長な方法
func generatePrimesV1(n int) []int {
	primes := []int{}

	for i := 2; i <= n; i++ {
		if isPrimeV1(i) {
			primes = append(primes, i)
		}
	}

	return primes
}

// cacheを使った方法
// 2から確認していき、素数であった数字の倍数は素数ではない数字としてキャッシュしておく
// (素数でない数字は、その手前の割り切れる素数によって、すべてそのあとはキャッシュされているので)
func generatePrimesV2(n int) []int {
	cache := make(map[int]struct{})
	primes := make([]int, 0)

	for i := 2; i <= n; i++ {
		_, isNotPrime := cache[i]
		if isNotPrime {
			continue
		}

		primes = append(primes, i)
		for j := i; j <= n; j += i {
			cache[j] = struct{}{}
		}
	}

	return primes
}

func generateSliceWithDefaultVal[T any](n int, val T) []T {
	result := make([]T, n)
	for i := range result {
		result[i] = val
	}
	return result
}

// エラトステネスの篩を使用する方法
func generatePrimesV3(n int) []int {
	if n <= 1 {
		return []int{}
	}

	cache := generateSliceWithDefaultVal(n, true)
	cache[0], cache[1] = false, false

	for current := 2; current < n; current++ {
		if !cache[current] {
			continue
		}

		i := 2
		j := current * 2
		for j < n {
			cache[j] = false
			i++
			j = current * i
		}
	}

	primes := []int{}
	for i := 2; i < len(cache); i++ {
		if cache[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

// -----------------------------
// 素数判定問題

// 冗長な方法
func isPrimeV1(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// √をつかった方法
func isPrimeV2(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	lessNum := int(math.Sqrt(float64(n))) + 1
	for i := 3; i < lessNum; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// √を使わずにループする方法
func isPrimeV3(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 || n%3 == 3 {
		return false
	}

	i := 3
	for i*i <= n {
		if n%i == 0 {
			return false
		}

		i++
	}

	return true
}
