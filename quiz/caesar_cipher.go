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

// alphabetの配列を使う方法
func caesarCipherV1(text string, shift int) (string, error) {
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

	return result, nil
}

// unicode pointを使う方法
func caesarCipherV2(text string, shift int) (string, error) {
	// s := ""
	for _, char := range text {
		if string(char) == " " {

		} else {

		}
	}

	return "", nil
}
