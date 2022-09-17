// l = [1, 3, 4, 2, 4, 5, 1, 6, 9, 8]  =>  l = [4, 2, 4, 6, 8, 1, 3, 5, 1, 9]
package quiz

import (
	"reflect"
	"testing"
)

func Test_orderEvenFirstOddLastV1(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{[]int{1, 3, 4, 2, 4, 5, 1, 6, 9, 8}},
			want: []int{4, 2, 4, 6, 8, 1, 3, 5, 1, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderEvenFirstOddLastV1(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("orderEvenFirstOddLastV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderEvenFirstOddLastV2(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{[]int{0, 1, 3, 4, 2, 4, 5, 1, 6, 9, 8}},
			want: []int{0, 8, 6, 4, 2, 4, 1, 5, 9, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderEvenFirstOddLastV2(tt.args.numbers)
			if !reflect.DeepEqual(tt.args.numbers, tt.want) {
				t.Errorf("got = %v, want %v", tt.args.numbers, tt.want)
			}
		})
	}
}
