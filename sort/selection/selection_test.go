package selection

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case1",
			args: args{[]int{8, 3, 4, 9, 1, 2}},
			want: []int{1, 2, 3, 4, 8, 9},
		},
		{
			name: "Case2",
			args: args{[]int{569, 236, 391, 867, 826, 461, 201, 338, 240, 248}},
			want: []int{201, 236, 240, 248, 338, 391, 461, 569, 826, 867},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("input = %v, Sort() = %v, want %v", tt.args.ints, got, tt.want)
			}
		})
	}
}
