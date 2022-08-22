package hanoi

import (
	"fmt"
	"testing"
)

func TestHanoiPrint(t *testing.T) {
	HanoiPrint(3, "A", "C", "B")
}

func TestHanoiList(t *testing.T) {
	actions := HanoiList(3, PositionA, PositionC, PositionB)

	for _, a := range actions {
		fmt.Printf("%+v\n", a)
	}
}
