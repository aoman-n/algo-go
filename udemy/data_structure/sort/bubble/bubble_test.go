package bubble

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
			args: args{ints: []int{8, 3, 5, 2, 7, 1}},
			want: []int{1, 2, 3, 5, 7, 8},
		},
		{
			name: "Case2",
			args: args{ints: []int{200, 100, 30, 40, 888, 999}},
			want: []int{30, 40, 100, 200, 888, 999},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
