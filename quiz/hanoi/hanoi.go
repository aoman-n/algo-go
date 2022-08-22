package hanoi

import "fmt"

// A, B, C

// diskNum: ディスクナンバー(下から1, 2, 3)
// stc:     元の位置
// dest:    目的の位置
// support: 中間値
func HanoiPrint(diskNum int, src string, dest string, support string) {
	if diskNum < 1 {
		return
	}

	HanoiPrint(diskNum-1, src, support, dest)
	fmt.Printf("move (%v), {%v} -> {%v}\n", diskNum, src, dest)
	HanoiPrint(diskNum-1, support, dest, src)
}

type Position string

// A, B, C
const (
	PositionA = "A"
	PositionB = "B"
	PositionC = "C"
)

type Action struct {
	diskNum int
	src     Position
	dest    Position
}

func HanoiList(diskNum int, src, dest, support Position) []Action {
	result := make([]Action, 0)

	var helper func(diskNum int, src, dest, support Position)
	helper = func(diskNum int, src, dest, support Position) {
		if diskNum < 1 {
			return
		}

		helper(diskNum-1, src, support, dest)
		result = append(result, Action{diskNum, src, dest})
		helper(diskNum-1, support, dest, src)
	}

	helper(diskNum, src, dest, support)

	return result
}
