package quiz

import (
	"reflect"
	"testing"
)

func Test_fermatLastTheorem(t *testing.T) {
	type args struct {
		maxNum    int
		squareNum int
	}
	tests := []struct {
		name string
		args args
		want [][3]int
		err  error
	}{
		{
			name: "Case1",
			args: args{
				maxNum:    10,
				squareNum: 2,
			},
			want: [][3]int{
				{3, 4, 5},
				{6, 8, 10},
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fermatLastTheorem(tt.args.maxNum, tt.args.squareNum)
			if err != tt.err {
				t.Fatalf("occurred unexpected error, got err = %v, but want err = %v", err, tt.err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fermatLastTheorem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isInteger(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{1.1},
			want: false,
		},
		{
			name: "case2",
			args: args{1.1111111111111111},
			want: false,
		},
		{
			name: "case3",
			args: args{1.0},
			want: true,
		},
		{
			name: "case4",
			args: args{222.000},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInteger(tt.args.n); got != tt.want {
				t.Errorf("isInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
