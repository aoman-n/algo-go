package quiz

import (
	"fmt"
	"testing"
)

func Test_vigenerChiperV1(t *testing.T) {
	text := "ATTACK SILICON VALLEY"
	vc := NewVigenerChiperV1(text, "CAT")

	encrypted := vc.Encrypt()

	vc2 := NewVigenerChiperV1(encrypted, "CAT")
	decrypted := vc2.Decrypt()

	fmt.Printf("text: %v\n", text)
	fmt.Printf("encrypted: %v\n", encrypted)
	fmt.Printf("decrypted: %v\n", decrypted)
}

func Test_vigenerChiperV2(t *testing.T) {
	text := "ATTACK SILICON VALLEY"

	vc := NewVigenerChiperV2(text, "CAT")
	encrypted := vc.Encrypt()

	vc2 := NewVigenerChiperV2(encrypted, "CAT")
	decrypted := vc2.Decrypt()

	fmt.Printf("text: %v\n", text)
	fmt.Printf("encrypted: %v\n", encrypted)
	fmt.Printf("decrypted: %v\n", decrypted)
}
