package quiz

import (
	"reflect"
	"testing"
)

// Input: 50 => Output: [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47]

func Test_generatePrimesV1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case1",
			args: args{50},
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePrimesV1(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generatePrimesV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generatePrimesV2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case1",
			args: args{50},
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePrimesV2(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generatePrimesV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
