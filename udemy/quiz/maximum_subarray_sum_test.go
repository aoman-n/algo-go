package quiz

import (
	"reflect"
	"testing"

	"github.com/laster18/algo-go/udemy/pkg/smath"
)

func TestGetMaxSequenceSumBad(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case1",
			args: args{
				numbers: []int{1, -2, 3, 6, -1, 2, 4, -5, 2},
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMaxSequenceSumBad(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMaxSequenceSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMaxSequenceSumGood(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case1",
			args: args{
				numbers: []int{1, -2, 3, 6, -1, 2, 4, -5, 2},
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMaxSequenceSum(tt.args.numbers, smath.MaxInt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMaxSequenceSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMaxSequenceSum(t *testing.T) {
	type args struct {
		numbers  []int
		operator func(int, int) int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMaxSequenceSum(tt.args.numbers, tt.args.operator); got != tt.want {
				t.Errorf("GetMaxSequenceSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMaxCircularSequenceSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case1",
			args: args{
				numbers: []int{1, -2, 3, 6, -1, 2, 4, -5, 2},
			},
			want: 15,
		},
		{
			name: "Case2",
			args: args{
				numbers: []int{-100, 1, -2, 3, 6, -1, 2, 4, -5, 2, -100},
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxCircularSequenceSum(tt.args.numbers); got != tt.want {
				t.Errorf("FindMaxCircularSequenceSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
