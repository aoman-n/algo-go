package insertion

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
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
			args: args{input: []int{8, 3, 4, 9, 1, 2}},
			want: []int{1, 2, 3, 4, 8, 9},
		},
		{
			name: "Case1",
			args: args{input: []int{681, 857, 76, 483, 602, 586}},
			want: []int{76, 483, 586, 602, 681, 857},
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
