package enigmaV2

import (
	"errors"
	"strings"
)

type Reflector struct {
	data map[string]string
}

var ErrUnMatch = errors.New("unmatched alphabet mapping")

func NewReflector(mapAlphabet string) (*Reflector, error) {
	r := new(Reflector)

	r.data = make(map[string]string)
	for i, charRune := range mapAlphabet {
		alpha := string(Alphabet[i])
		r.data[alpha] = string(charRune)
	}

	for key, val := range r.data {
		if r.data[val] != key {
			return nil, ErrUnMatch
		}
	}

	return r, nil
}

func (r *Reflector) Reflect(index int) int {
	char := string(Alphabet[index])
	char = r.data[char]
	return strings.Index(Alphabet, char)
}
