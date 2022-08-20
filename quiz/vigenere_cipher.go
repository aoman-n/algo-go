package quiz

import (
	"strings"
	"unicode"
)

// a-z: 26
// var uppercaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// var lowercaseAlphabet = "abcdefghijklmnopqrstuvwxyz"

func getIndexFromRune(r rune) int {
	if unicode.IsUpper(r) {
		return strings.IndexRune(uppercaseAlphabet, r)
	} else {
		return strings.IndexRune(lowercaseAlphabet, r)
	}
}

func generateKey(text, keyword string) string {
	key := keyword
	remainLength := len(text) - len(keyword)

	if remainLength < 0 {
		return ""
	}

	for i := range make([]struct{}, remainLength) {
		key += string(keyword[i%len(keyword)])
	}

	return key
}

// @see: https://ja.wikipedia.org/wiki/%E3%83%B4%E3%82%A3%E3%82%B8%E3%83%A5%E3%83%8D%E3%83%AB%E6%9A%97%E5%8F%B7

type VigenerChiperV1 struct {
	val string
	key string
}

func NewVigenerChiperV1(val, keyword string) *VigenerChiperV1 {
	v := new(VigenerChiperV1)
	v.val = val
	v.key = generateKey(val, keyword)
	return v
}

func (v *VigenerChiperV1) Encrypt() string {
	alphabetLen := len(uppercaseAlphabet)
	result := ""

	for i, charRune := range v.val {
		newIndex := (getIndexFromRune(charRune) + getIndexFromRune(rune(v.key[i]))) % alphabetLen

		if unicode.IsUpper(charRune) {
			result += string(uppercaseAlphabet[newIndex])
		} else if unicode.IsLower(charRune) {
			result += string(lowercaseAlphabet[newIndex])
		} else {
			result += string(charRune)
		}
	}

	return result
}

func (v *VigenerChiperV1) Decrypt() string {
	alphabetLen := len(uppercaseAlphabet)
	result := ""

	for i, charRune := range v.val {
		newIndex := (getIndexFromRune(charRune) - getIndexFromRune(rune(v.key[i])) + alphabetLen) % alphabetLen

		if unicode.IsUpper(charRune) {
			result += string(uppercaseAlphabet[newIndex])
		} else if unicode.IsLower(charRune) {
			result += string(lowercaseAlphabet[newIndex])
		} else {
			result += string(charRune)
		}
	}

	return result
}

// VigenerChiperV2
// unicode pointを使用した暗号化
// アルファベットは大文字を使用
type VigenerChiperV2 struct {
	val string
	key string
}

func NewVigenerChiperV2(val, keyword string) *VigenerChiperV2 {
	v := new(VigenerChiperV2)
	v.val = val
	v.key = generateKey(val, keyword)
	return v
}

func (v *VigenerChiperV2) Encrypt() string {
	alphabetLen := len(uppercaseAlphabet)
	result := ""

	for i, charRune := range v.val {
		if unicode.IsSpace(charRune) {
			result += string(charRune)
			continue
		}

		keyRune := rune(v.key[i])
		newRune := (charRune+keyRune)%rune(alphabetLen) + 'A'
		result += string(newRune)
	}

	return result
}

func (v *VigenerChiperV2) Decrypt() string {
	alphabetLen := len(uppercaseAlphabet)
	result := ""

	for i, charRune := range v.val {
		if unicode.IsSpace(charRune) {
			result += string(charRune)
			continue
		}

		keyRune := rune(v.key[i])
		newRune := (charRune-keyRune+rune(alphabetLen))%rune(alphabetLen) + 'A'
		result += string(newRune)
	}

	return result
}

func VegenereEncrypt(text, keyword string) string {

	return ""
}

func VegenereDecrypt(text, keyword string) string {
	return ""
}
