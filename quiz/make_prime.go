package quiz

func isPrime(n int) bool {
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

// 冗長な方法
func generatePrimesV1(n int) []int {
	primes := []int{}

	for i := 2; i <= n; i++ {
		if isPrime(i) {
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
