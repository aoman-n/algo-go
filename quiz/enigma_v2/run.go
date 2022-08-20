package enigmaV2

import (
	"fmt"
	"math/rand"
)

func Run() error {
	// create plugboard
	plugboard := NewPlugboard(getRandAlphabets())

	// create rotors
	rotor1 := NewRotor(getRandAlphabets(), 3)
	rotor2 := NewRotor(getRandAlphabets(), 2)
	rotor3 := NewRotor(getRandAlphabets(), 1)

	// create reflector
	reflector, err := NewReflector(getReflectAlphabets())
	if err != nil {
		return err
	}

	// create enigma
	enigma := NewEnigma(plugboard, []*Rotor{rotor1, rotor2, rotor3}, reflector)

	origin := "ATTACK SILICON VALLEY by engineer"
	encrypted := enigma.Encrypt(origin)
	decrypted := enigma.Decrypt(encrypted)

	fmt.Printf("origin: %v\n", origin)
	fmt.Printf("encrypted: %v\n", encrypted)
	fmt.Printf("decrypted: %v\n", decrypted)

	return nil
}

func getRandAlphabets() string {
	alphabetRunes := []rune(Alphabet)

	rand.Shuffle(len(alphabetRunes), func(i, j int) {
		alphabetRunes[i], alphabetRunes[j] = alphabetRunes[j], alphabetRunes[i]
	})

	return string(alphabetRunes)
}

func getReflectAlphabets() string {
	indexes := make([]int, 0, len(Alphabet))
	for i := range Alphabet {
		indexes = append(indexes, i)
	}

	pairs := [][2]int{}

	for len(indexes) > 0 {
		pairs = append(pairs, [2]int{randPop(&indexes), randPop(&indexes)})
	}

	alphabets := []rune(Alphabet)
	for _, pair := range pairs {
		x, y := pair[0], pair[1]
		alphabets[x], alphabets[y] = alphabets[y], alphabets[x]
	}

	return string(alphabets)
}

// [0, 1, 2, 3, ..., 25]

// 012345
// ABCDEF

// "ABCDEF"
// "CDABFE"

// [0, 2], [1, 3], [4, 5]

func randPop(list *[]int) int {
	if len(*list) == 1 {
		tmp := (*list)[0]
		*list = []int{}
		return tmp
	}

	randomIndex := rand.Intn(len(*list))

	tmp := (*list)[randomIndex]

	newList := make([]int, len(*list)-1)
	copy(newList, *list)
	copy(newList[randomIndex:], (*list)[randomIndex+1:])
	*list = newList

	return tmp
}
