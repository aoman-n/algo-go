package enigmaV2

import "strings"

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

func (p *Plugboard) Forward(index int) int {
	char := string(p.alphabet[index])
	char = p.forwardMap[char]
	return strings.Index(p.alphabet, char)
}

func (p *Plugboard) Backward(index int) int {
	char := string(p.alphabet[index])
	char = p.backwardMap[char]
	return strings.Index(p.alphabet, char)
}
