package quiz

import (
	"testing"

	"github.com/aoman-n/algo-go/udemy/pkg/slice"
)

func Test_phoneMnemonicV1(t *testing.T) {
	got, _ := phoneMnemonicV1("568-379-8466")
	if !slice.Exists(got, "LOVEPYTHON") {
		t.Errorf(`not found "LOVEPYTHON"`)
	}
}

func Test_phoneMnemonicV2(t *testing.T) {
	got, _ := phoneMnemonicV2("568-379-8466")
	if !slice.Exists(got, "LOVEPYTHON") {
		t.Errorf(`not found "LOVEPYTHON"`)
	}
}
