package quiz

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			got := AllPermutation(tt.args.list)
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}
