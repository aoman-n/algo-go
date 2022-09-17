package quiz

import (
	"reflect"
	"testing"
)

func TestFindSymmetricPairs(t *testing.T) {
	type args struct {
		pairs []Pair
	}
	tests := []struct {
		name string
		args args
		want []Pair
	}{
		{
			name: "Case1",
			args: args{
				pairs: []Pair{{1, 2}, {3, 5}, {4, 7}, {5, 3}, {7, 4}},
			},
			want: []Pair{{5, 3}, {7, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSymmetricPairs(tt.args.pairs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSymmetricPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
