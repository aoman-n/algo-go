package quiz

import (
	"errors"
	"strings"
	"unicode"
)

// a-z: 26
var uppercaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var lowercaseAlphabet = "abcdefghijklmnopqrstuvwxyz"

var ErrInvalidText = errors.New("invalid text error")

// alphabetの文字列を使う方法
func caesarCipherV1(text string, shift int) string {
	result := ""

	for _, char := range text {
		alphabet := ""

		switch {
		case unicode.IsUpper(char):
			alphabet = uppercaseAlphabet
		case unicode.IsLower(char):
			alphabet = lowercaseAlphabet
		default:
			result += string(char)
			continue
		}

		alphaIndex := strings.IndexRune(alphabet, char)

		i := alphaIndex + shift
		var newIndex int
		if i < 0 {
			newIndex = len(alphabet) + i
		} else {
			newIndex = (alphaIndex + shift) % len(alphabet)
		}
		result += string(alphabet[newIndex])

	}

	return result
}

// rune(code point)
// a-z: 97-122
// A-Z: 65-90

// unicode pointを使う方法
func caesarCipherV2(text string, shift int) string {
	result := ""
	lenAlphabet := 'Z' - 'A' + 1

	for _, charRune := range text {
		if unicode.IsUpper(charRune) {
			newCharRune := (charRune+rune(shift)-'A')%lenAlphabet + 'A'
			if newCharRune < 'A' {
				result += string(newCharRune + rune(lenAlphabet))
			} else {
				result += string(newCharRune)
			}
		} else if unicode.IsLower(charRune) {
			newCharRune := (charRune+rune(shift)-'a')%lenAlphabet + 'a'
			if newCharRune < 'a' {
				result += string(newCharRune + rune(lenAlphabet))
			} else {
				result += string(newCharRune)
			}
		} else {
			result += string(charRune)
		}
	}

	return result
}

// alphabetの文字列を使う方法
func caesarCipherHackV1(text string) []string {
	lenAlphabet := len(uppercaseAlphabet)
	result := make([]string, 0, lenAlphabet)

	for i := range make([]struct{}, lenAlphabet) {
		incrementCount := i + 1
		tmp := ""

		for _, charRune := range text {
			var alphabet string
			if unicode.IsUpper(charRune) {
				alphabet = uppercaseAlphabet
			} else if unicode.IsLower(charRune) {
				alphabet = lowercaseAlphabet
			} else {
				tmp += string(charRune)
				continue
			}

			index := strings.IndexRune(alphabet, charRune)
			newIndex := (index + incrementCount) % len(alphabet)
			tmp += string(alphabet[newIndex])
		}

		result = append(result, tmp)
	}

	return result
}

// unicode pointを使う方法
func caesarCipherHackV2(text string) []string {
	lenAlphabet := int('Z' - 'A' + 1)
	result := make([]string, 0, lenAlphabet)

	for i := 1; i <= lenAlphabet; i++ {
		tmp := ""

		for _, charRune := range text {
			if unicode.IsUpper(charRune) {
				newRune := (charRune-'A'+rune(i))%rune(lenAlphabet) + 'A'
				tmp += string(newRune)
			} else if unicode.IsLower(charRune) {
				newRune := (charRune-'a'+rune(i))%rune(lenAlphabet) + 'a'
				tmp += string(newRune)
			} else {
				tmp += string(charRune)
			}
		}

		result = append(result, tmp)
	}

	return result
}
