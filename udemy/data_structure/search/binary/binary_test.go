package binary

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	type args struct {
		s []int
		v int
	}
	tests := []struct {
		name string
		args args
		want IndexInt
	}{
		{
			name: "Case1",
			args: args{s: []int{1, 2, 3, 4, 5, 6}, v: 3},
			want: 2,
		},
		{
			name: "Case2",
			args: args{s: []int{1, 2, 3, 4, 5, 6}, v: 10},
			want: -1,
		},
		{
			name: "Case2",
			args: args{s: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, v: 10},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.s, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
