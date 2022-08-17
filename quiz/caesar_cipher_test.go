package quiz

import (
	"fmt"
	"testing"
)

func Test_caesarCipherV1(t *testing.T) {
	text := "ATTACK SILICON VALLEY by engineer"

	ret, err := caesarCipherV1(text, 3)
	if err != nil {
		t.Fatalf("unexpected error, err: %v", err)
	}

	fmt.Printf("ret: %v\n", ret)

	ret2, _ := caesarCipherV1(ret, -3)

	fmt.Printf("ret2: %v\n", ret2)
}
