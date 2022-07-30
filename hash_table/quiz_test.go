package hashtable

import (
	"testing"
)

func TestGetPair(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 bool
	}{
		{
			name: "Case1",
			args: args{
				numbers: []int{11, 2, 5, 9, 10, 3},
				target:  12,
			},
			want:  2,
			want1: 10,
			want2: true,
		},
		{
			name: "Case2",
			args: args{
				numbers: []int{11, 2, 5, 8, 10, 3},
				target:  20,
			},
			want:  0,
			want1: 0,
			want2: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := GetPair(tt.args.numbers, tt.args.target)
			if got != tt.want {
				t.Errorf("GetPair() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetPair() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("GetPair() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestGetPairHalfSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 bool
	}{
		{
			name: "Case1",
			args: args{
				numbers: []int{11, 2, 5, 9, 10, 3},
			},
			want:  11,
			want1: 9,
			want2: true,
		},
		{
			name: "Case2",
			args: args{
				numbers: []int{11, 2, 5, 8, 11, 3},
			},
			want:  0,
			want1: 0,
			want2: false,
		},
		{
			name: "Case3",
			args: args{
				numbers: []int{11, 2, 14, 11, 3, 9},
			},
			want:  11,
			want1: 14,
			want2: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := GetPairHalfSum(tt.args.numbers)
			if got != tt.want {
				t.Errorf("GetPair() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetPair() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("GetPair() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
