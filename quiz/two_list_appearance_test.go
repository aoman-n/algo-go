package quiz

import (
	"reflect"
	"testing"
)

func TestMinCountRemove(t *testing.T) {
	type args struct {
		x []int
		y []int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		{
			name: "Case1",
			args: args{
				x: []int{1, 2, 3, 4, 4, 5, 5, 8, 10},
				y: []int{4, 5, 5, 5, 6, 7, 8, 8, 10},
			},
			want:  []int{1, 2, 3, 4, 4, 10},
			want1: []int{5, 5, 5, 6, 7, 8, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MinCountRemove(tt.args.x, tt.args.y)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinCountRemove() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("MinCountRemove() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
