package quiz

import (
	"unicode"
)

// input: 'This is a pen. This is an apple. Applepen.'

func countChars(chars string) (string, int) {
	m := make(map[string]int)
	mostCount := -1
	mostStr := ""
	for _, c := range chars {
		if unicode.IsSpace(c) {
			continue
		}

		s := string(unicode.ToLower(c))
		m[s] += 1
		count := m[s]
		if count > mostCount {
			mostStr = s
			mostCount = count
		}
	}

	return mostStr, mostCount
}
