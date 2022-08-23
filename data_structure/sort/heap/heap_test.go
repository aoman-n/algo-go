package heap

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/laster18/algo-go/pkg/generator"
)

func TestHeap(t *testing.T) {
	h := BuildMaxHeap([]int{4, 2, 5, 2, 6, 10, 3, 7, 6, 5})
	fmt.Println(h.List())
}

func TestSort(t *testing.T) {
	fmt.Println(generator.Ints(10))
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case1",
			args: args{input: []int{1, 8, 3, 9, 4, 5, 7}},
			want: []int{1, 3, 4, 5, 7, 8, 9},
		},
		{
			name: "Case1",
			args: args{input: []int{904, 689, 942, 935, 970, 3, 9, 951, 222, 532}},
			want: []int{3, 9, 222, 532, 689, 904, 935, 942, 951, 970},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("input = %v, Sort() = %v, want %v", tt.args.input, got, tt.want)
			}
		})
	}
}
