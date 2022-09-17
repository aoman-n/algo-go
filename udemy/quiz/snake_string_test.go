package quiz

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestSnakeStringV1(t *testing.T) {
	input := ""
	for range make([]struct{}, 5) {
		for j := range make([]struct{}, 10) {
			input += strconv.Itoa(j)
		}
	}

	got := SnakeStringV1(input)

	for _, val := range got {
		fmt.Println(strings.Join(val, ""))
	}
}

func TestSnakeStringV2(t *testing.T) {
	input := ""
	for range make([]struct{}, 5) {
		for j := range make([]struct{}, 10) {
			input += strconv.Itoa(j)
		}
	}

	got := SnakeStringV2(input, 6)

	for _, val := range got {
		fmt.Println(strings.Join(val, ""))
	}
}
