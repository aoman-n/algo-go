package quiz

func isPalindrome(str string) bool {
	l := len(str)
	if l <= 0 {
		return false
	}
	if l <= 1 {
		return true
	}

	start, end := 0, l-1

	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}

	return true
}

func findPalindrome(str string, left, right int) []string {
	result := []string{}

	for left >= 0 && right <= len(str)-1 {
		if str[left] != str[right] {
			break
		}

		result = append(result, str[left:right+1])

		left--
		right++
	}

	return result
}

func findAllPalindrome(str string) []string {
	l := len(str)
	if l == 0 {
		return []string{}
	}
	if l <= 1 {
		return []string{str}
	}

	result := make([]string, 0)

	for i := 1; i < l-1; i++ {
		result = append(result, findPalindrome(str, i-1, i+1)...)
		result = append(result, findPalindrome(str, i-1, i)...)
	}

	return result
}
