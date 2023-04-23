package tips_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aoman-n/algo-go/container/tips"
)

func TestGetMax(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			input: []int{31, 51, 59, 46},
			want:  59,
		},
		{
			input: []int{200, 100, 50, 203, 130},
			want:  203,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("input(%v)", tt.input), func(t *testing.T) {
			t.Parallel()

			got := tips.GetMax(tt.input)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tips.GetMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
