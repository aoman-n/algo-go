package enigmav1

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Plugboard(t *testing.T) {
	p := NewPlugboard("BADC")
	fmt.Printf("p: %+v\n", p)

	forwarded := p.Forward(strings.Index(Alphabet, "A"))
	backwarded := p.Backward(forwarded)
	ret := string(Alphabet[backwarded])
	fmt.Printf("ret: %v\n", ret)
}

func Test_Rotor(t *testing.T) {
	origin := "A"

	p := NewRotor("BADC", 3)
	forwarded := p.Forward(strings.Index(Alphabet, origin))
	backwarded := p.Backward(forwarded)
	fmt.Printf("origin: %v\n", origin)
	fmt.Printf("forwarded: %v\n", string(Alphabet[forwarded]))
	fmt.Printf("backwarded: %v\n", string(Alphabet[backwarded]))

	p.Rotate()

	forwarded2 := p.Forward(strings.Index(Alphabet, origin))
	backwarded2 := p.Backward(forwarded2)
	fmt.Printf("origin2: %v\n", origin)
	fmt.Printf("forwarded2: %v\n", string(Alphabet[forwarded2]))
	fmt.Printf("backwarded2: %v\n", string(Alphabet[backwarded2]))
}

func TestReflector(t *testing.T) {
	r, err := NewReflector("BADC")
	if err != nil {
		t.Errorf("occurred unexpected error: %v", err)
	}

	origin := "A"
	fmt.Printf("origin: %v\n", origin)

	reflected := r.Reflect(strings.Index(Alphabet, origin))
	reflectedChar := string(Alphabet[reflected])
	fmt.Printf("outChar: %s\n", reflectedChar)

	restored := r.Reflect(reflected)
	restoredChar := string(Alphabet[restored])
	fmt.Printf("restoredChar: %s\n", restoredChar)
}

func TestEnigma(t *testing.T) {
	// create plugboard
	plugboard := NewPlugboard(GetRandomAlphabet())

	// create rotors
	rotor1 := NewRotor(GetRandomAlphabet(), 3)
	rotor2 := NewRotor(GetRandomAlphabet(), 2)
	rotor3 := NewRotor(GetRandomAlphabet(), 1)

	// create reflector
	reflector, err := ReflectorFactory()
	assert.NoError(t, err)

	// create enigma
	enigma := NewEnigma(plugboard, []*Rotor{rotor1, rotor2, rotor3}, reflector)

	encrypted := enigma.Encrypt("ATTACK SILICON VALLEY by engineer")
	decrypted := enigma.Decrypt(encrypted)

	fmt.Printf("encrypted: %v\n", encrypted)
	fmt.Printf("decrypted: %v\n", decrypted)
}
