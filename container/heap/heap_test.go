package heap_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aoman-n/algo-go/container/heap"
)

func intCmp(x, y int) int {
	if x < y {
		return -1
	}
	if x == y {
		return 0
	}

	return 1
}

func TestHeap_BuildMaxHeap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{1, 2, 3},
			want:  []int{3, 2, 1},
		},
		{
			input: []int{-3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  []int{10, 7, 9, 5, 6, 8, 3, 4, 0, -2, 1, -3, 2, -1},
		},
		{
			input: []int{7, 8, 2, 3, 1, 4, 3, 2},
			want:  []int{8, 7, 4, 3, 1, 2, 3, 2},
		},
		{
			input: []int{8, 4, 13, 10, 18},
			want:  []int{18, 10, 13, 8, 4},
		},
		{
			input: []int{3, 100, 201, 56, 8, 591, 985, 291},
			want:  []int{985, 291, 591, 100, 8, 3, 201, 56},
		},
		{
			input: []int{879, 487, 98, 397, 610, 150, 474, 977, 404, 478, 623, 554, 306},
			want:  []int{977, 879, 554, 487, 623, 306, 474, 397, 404, 478, 610, 150, 98},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(fmt.Sprintf("input(%+v) want(%+v)", tt.input, tt.want), func(t *testing.T) {
			got := heap.BuildMaxHeap(tt.input, intCmp)

			if !reflect.DeepEqual(got.List(), tt.want) {
				t.Errorf("BuildMaxHeap() = %v, but want %v", got.List(), tt.want)
			}
		})
	}
}

func TestHeap_PriorityQueue(t *testing.T) {
	intHeap := heap.BuildMaxHeap([]int{1, 2, 3}, intCmp)

	fmt.Println(intHeap.Pop()) // 3
	intHeap.Insert(-1)
	intHeap.Insert(10)
	intHeap.Insert(-2)
	fmt.Println(intHeap.Pop()) // 10
	fmt.Println(intHeap.Pop()) // 2
	fmt.Println(intHeap.Top()) // 1
	fmt.Println(intHeap.Pop()) // 1
	fmt.Println(intHeap.Pop()) // -1
	fmt.Println(intHeap.Pop()) // -2
}

func TestHeap_Sort(t *testing.T) {}
