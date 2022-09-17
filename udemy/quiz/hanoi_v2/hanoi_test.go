package hanoi

import (
	"fmt"
	"testing"
)

func TestHanoiPrint(t *testing.T) {
	HanoiPrint(3, PositionA, PositionC, PositionB)
}

func TestHanoi(t *testing.T) {
	actions := Hanoi(3, PositionA, PositionC, PositionB)
	for _, a := range actions {
		fmt.Println(a)
	}
}
