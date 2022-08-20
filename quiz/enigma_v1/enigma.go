package enigmav1

import (
	"fmt"
	"math/rand"
	"strings"
	"unicode"
)

const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Plugboard struct {
	alphabet    string
	forwardMap  map[string]string
	backwardMap map[string]string
}

func NewPlugboard(mapAlphabet string) *Plugboard {
	p := new(Plugboard)
	p.alphabet = Alphabet
	p.mapping(mapAlphabet)
	return p
}

func (p *Plugboard) mapping(mapAlphabet string) {
	p.forwardMap = make(map[string]string)
	p.backwardMap = make(map[string]string)
	for i, r := range mapAlphabet {
		p.forwardMap[string(p.alphabet[i])] = string(r)
		p.backwardMap[string(r)] = string(p.alphabet[i])
	}
}

func (p *Plugboard) Forward(i int) int {
	char := string(p.alphabet[i])
	char = p.forwardMap[char]
	return strings.Index(p.alphabet, char)
}

func (p *Plugboard) Backward(i int) int {
	char := string(p.alphabet[i])
	char = p.backwardMap[char]
	return strings.Index(p.alphabet, char)
}

type Rotor struct {
	*Plugboard
	rotations int
	offset    int
}

func NewRotor(mapAlphabet string, offset int) *Rotor {
	r := new(Rotor)
	r.Plugboard = NewPlugboard(mapAlphabet)
	r.rotations = 0
	r.offset = offset
	return r
}

func (r *Rotor) Rotate() int {
	r.alphabet = r.alphabet[r.offset:] + r.alphabet[:r.offset]
	r.rotations += r.offset
	return r.rotations
}

func (r *Rotor) Reset() {
	r.alphabet = Alphabet
	r.rotations = 0
}

type Reflector struct {
	mapAlphabet map[string]string
}

func NewReflector(mapAlphabet string) (*Reflector, error) {
	r := new(Reflector)

	m := make(map[string]string)
	for i, s := range mapAlphabet {
		m[string(s)] = string(Alphabet[i])
	}
	r.mapAlphabet = m

	for x, y := range r.mapAlphabet {
		if r.mapAlphabet[y] != x {
			return nil, fmt.Errorf("not matched values")
		}
	}

	return r, nil
}

func (r *Reflector) Reflect(i int) int {
	char := string(Alphabet[i])
	char = string(r.mapAlphabet[char])
	return strings.Index(Alphabet, char)
}

type Enigma struct {
	plugboard *Plugboard
	rotors    []*Rotor
	reflector *Reflector
}

func NewEnigma(plugboard *Plugboard, rotors []*Rotor, reflector *Reflector) *Enigma {
	return &Enigma{
		plugboard: plugboard,
		rotors:    rotors,
		reflector: reflector,
	}
}

func (e *Enigma) Encrypt(text string) string {
	result := ""
	for _, charRune := range text {
		result += e.goThrough(charRune)
	}

	return result
}

func (e *Enigma) Decrypt(text string) string {
	for _, r := range e.rotors {
		r.Reset()
	}

	result := ""
	for _, charRune := range text {
		result += e.goThrough(charRune)
	}

	return result
}

func (e *Enigma) goThrough(r rune) string {
	if unicode.IsSpace(r) {
		return " "
	}

	r = unicode.ToUpper(r)
	indexNum := strings.IndexRune(Alphabet, r)

	indexNum = e.plugboard.Forward(indexNum)
	for _, r := range e.rotors {
		indexNum = r.Forward(indexNum)
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

func GetRandomAlphabet() string {
	alphabets := make([]string, 0, len(Alphabet))
	for _, r := range Alphabet {
		alphabets = append(alphabets, string(r))
	}

	rand.Shuffle(len(alphabets), func(i, j int) {
		alphabets[i], alphabets[j] = alphabets[j], alphabets[i]
	})

	result := ""
	for _, s := range alphabets {
		result += s
	}

	return result
}

func ReflectorFactory() (*Reflector, error) {
	return NewReflector(getReflectorAlphabets())
}

func getReflectorAlphabets() string {
	indexes := make([]int, 0, len(Alphabet))
	for i := range Alphabet {
		indexes = append(indexes, i)
	}

	pairs := [][2]int{}
	for len(indexes) > 0 {
		pairs = append(pairs, [2]int{randPop(&indexes), randPop(&indexes)})
	}

	alphabetRunes := []rune(Alphabet)
	for _, v := range pairs {
		x, y := v[0], v[1]
		alphabetRunes[x], alphabetRunes[y] = alphabetRunes[y], alphabetRunes[x]
	}

	return string(alphabetRunes)
}

func randPop(ints *[]int) int {
	l := len(*ints)

	if l == 1 {
		tmp := (*ints)[0]
		*ints = []int{}
		return tmp
	}

	index := rand.Intn(l - 1)
	newInts := make([]int, l-1)
	poped := (*ints)[index]

	copy(newInts[:index], (*ints)[:index])
	copy(newInts[index:], (*ints)[index+1:])

	*ints = newInts

	return poped
}
