package enigmaV2

type Rotor struct {
	*Plugboard
	// offset rotateの基準値
	offset int
	// rotations rotateと合計値
	rotations int
}

func NewRotor(mapAlphabet string, offset int) *Rotor {
	r := new(Rotor)
	r.Plugboard = NewPlugboard(mapAlphabet)
	r.offset = offset
	r.rotations = 0
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
