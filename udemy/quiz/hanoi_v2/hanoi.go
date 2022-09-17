package hanoi

import "fmt"

type Position string

const (
	PositionA Position = "A"
	PositionB Position = "B"
	PositionC Position = "C"
)

func HanoiPrint(disk int, src, dest, support Position) {
	if disk < 1 {
		return
	}

	HanoiPrint(disk-1, src, support, dest)
	fmt.Printf("disk is {%v}, {%v} -> {%v}\n", disk, src, dest)
	HanoiPrint(disk-1, support, dest, src)
}

type Action struct {
	Disk int
	From Position
	To   Position
}

func Hanoi(disk int, src, dest, support Position) []Action {
	actions := make([]Action, 0)

	var helper func(disk int, src, dest, support Position)

	helper = func(disk int, src, dest, support Position) {
		if disk < 1 {
			return
		}

		helper(disk-1, src, support, dest)
		actions = append(actions, Action{disk, src, dest})
		helper(disk-1, support, dest, src)
	}

	helper(disk, src, dest, support)

	return actions
}
