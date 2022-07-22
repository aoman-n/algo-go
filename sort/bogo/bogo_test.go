package bogo

import (
	"testing"
)

func Test_Run(t *testing.T) {
	numbers := randomInts(10)
	t.Logf("before numbers: %v\n", numbers)
	Sort(numbers)
	t.Logf("after numbers: %v\n", numbers)
}
