package quiz

import (
	"reflect"
	"testing"
)

func TestAllPermutation(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Case1",
			args: args{
				list: []int{1, 2, 3},
			},
			want: [][]int{
				{1, 2, 3},
				{2, 1, 3},
				{2, 3, 1},
				{1, 3, 2},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllPermutation(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swapHeadTo(t *testing.T) {
	type args struct {
		list []int
		i    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case1",
			args: args{
				list: []int{111, 1, 2, 3, 4, 5},
				i:    2,
			},
			want: []int{1, 2, 111, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swapHeadTo(tt.args.list, tt.args.i)
			t.Logf("got = %v, want = %v", tt.args.list, tt.want)
			if !reflect.DeepEqual(tt.args.list, tt.want) {
				t.Errorf("got = %v, want = %v", tt.args.list, tt.want)
			}
		})
	}
}
