package quiz

import (
	"fmt"
	"testing"
)

func Test_caesarCipherV1(t *testing.T) {
	text := "ATTACK SILICON VALLEY by engineer"

	ret := caesarCipherV1(text, 3)
	fmt.Printf("ret: %v\n", ret)

	ret2 := caesarCipherV1(ret, -3)
	fmt.Printf("ret2: %v\n", ret2)
}

func Test_caesarCipherV2(t *testing.T) {
	text := "ATTACK SILICON VALLEY by engineer"

	ret := caesarCipherV2(text, 3)
	fmt.Printf("ret: %v\n", ret)

	ret2 := caesarCipherV2(ret, -3)
	fmt.Printf("ret2: %v\n", ret2)
}

func Test_caesarCipherHackV1(t *testing.T) {
	ret := caesarCipherHackV1("DWWDFN VLOLFRQ YDOOHB eb hqjlqhhu")

	for i, t := range ret {
		fmt.Printf("[%d] %s\n", i, t)
	}
}

func Test_caesarCipherHackV2(t *testing.T) {
	ret := caesarCipherHackV2("DWWDFN VLOLFRQ YDOOHB eb hqjlqhhu")

	for i, t := range ret {
		fmt.Printf("[%d] %s\n", i, t)
	}
}
