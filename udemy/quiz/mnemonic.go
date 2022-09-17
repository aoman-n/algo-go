package quiz

import (
	"fmt"
	"strconv"
	"strings"
)

func getNumbers(chars string) ([]int, error) {
	result := make([]int, 0, len(chars))
	for _, char := range chars {
		n, err := strconv.Atoi(string(char))
		if err != nil {
			return []int{}, err
		}
		result = append(result, n)
	}

	return result, nil
}

var numAlphabetMapping = map[int]string{
	0: "+",
	1: "@",
	2: "ABC",
	3: "DEF",
	4: "GHI",
	5: "JKL",
	6: "MNO",
	7: "PQRS",
	8: "TUV",
	9: "WXYZ",
}

// recursiveを使う方法
func phoneMnemonicV1(phoneNumber string) ([]string, error) {
	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")
	phoneNumbers, err := getNumbers(phoneNumber)
	if err != nil {
		return nil, err
	}

	candidate := []string{}
	tmp := make([]string, len(phoneNumber))

	var findCandidateAlphabet func(digit int)
	findCandidateAlphabet = func(digit int) {
		if digit == len(tmp) {
			candidate = append(candidate, strings.Join(tmp, ""))
		} else {
			targetNum := phoneNumbers[digit]
			for _, char := range numAlphabetMapping[targetNum] {
				tmp[digit] = string(char)
				findCandidateAlphabet(digit + 1)
			}
		}
	}

	findCandidateAlphabet(0)

	return candidate, nil
}

// '23'

// type node[T any] struct {
// 	val  T
// 	next *node[T]
// }

// func newNode[T any](val T) *node[T] {
// 	return &node[T]{
// 		val:  val,
// 		next: nil,
// 	}
// }

type stack struct {
	data []string
}

func newStack() *stack {
	return &stack{
		data: []string{},
	}
}

func (s *stack) pop() string {
	if len(s.data) <= 0 {
		return ""
	}

	ret := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return ret
}

func (s *stack) push(str string) {
	s.data = append(s.data, str)
}

func (s *stack) len() int {
	l := len(s.data)
	return l
}

// stackを使う方法
func phoneMnemonicV2(phoneNumber string) ([]string, error) {
	phoneNumbers, err := getNumbers(strings.ReplaceAll(phoneNumber, "-", ""))
	if err != nil {
		return nil, err
	}

	s := newStack()
	s.push("")
	candidate := []string{}

	for s.len() != 0 {
		str := s.pop()
		if len(str) == len(phoneNumbers) {
			candidate = append(candidate, str)
		} else {
			currentAlphabets := numAlphabetMapping[phoneNumbers[len(str)]]
			for _, alpha := range currentAlphabets {
				s.push(str + string(alpha))
			}
		}
	}

	fmt.Printf("candidate: %v\n", candidate)

	return candidate, nil
}
