package heap

import (
	"reflect"
	"testing"
)

func TestTopNWithHeap(t *testing.T) {
	type args struct {
		list []string
		n    int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Case1",
			args: args{
				list: []string{"python", "java", "c", "go", "python", "c", "go", "python"},
				n:    3,
			},
			want: []string{"python", "go", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopNWithHeap(tt.args.list, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopNWithHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}
