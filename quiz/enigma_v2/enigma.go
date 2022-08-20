package enigmaV2

import (
	"strings"
	"unicode"
)

const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Enigma struct {
	plugboard *Plugboard
	rotors    []*Rotor
	reflector *Reflector
}

func NewEnigma(plugboard *Plugboard, rotors []*Rotor, reflector *Reflector) *Enigma {
	e := new(Enigma)
	e.plugboard = plugboard
	e.rotors = rotors
	e.reflector = reflector
	return e
}

func (e *Enigma) Encrypt(text string) string {
	result := ""
	for _, charRune := range text {
		result += e.goThrough(charRune)
	}
	return result
}

func (e *Enigma) Decrypt(text string) string {
	for _, rotor := range e.rotors {
		rotor.Reset()
	}

	result := ""
	for _, charRune := range text {
		result += e.goThrough(charRune)
	}
	return result
}

func (e *Enigma) goThrough(charRune rune) string {
	if unicode.IsSpace(charRune) {
		return " "
	}

	charRune = unicode.ToUpper(charRune)
	indexNum := strings.IndexRune(Alphabet, charRune)

	indexNum = e.plugboard.Forward(indexNum)

	for _, rotor := range e.rotors {
		indexNum = rotor.Forward(indexNum)
	}

	indexNum = e.reflector.Reflect(indexNum)

	for i := len(e.rotors) - 1; i >= 0; i-- {
		indexNum = e.rotors[i].Backward(indexNum)
	}

	indexNum = e.plugboard.Backward(indexNum)

	resultChar := string(Alphabet[indexNum])

	for i := len(e.rotors) - 1; i >= 0; i-- {
		if e.rotors[i].Rotate()%len(Alphabet) != 0 {
			break
		}
	}

	return resultChar
}
