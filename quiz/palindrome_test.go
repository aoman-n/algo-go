package quiz

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case1",
			args: args{str: "cabac"},
			want: true,
		},
		{
			name: "Case2",
			args: args{str: "abba"},
			want: true,
		},
		{
			name: "Case2",
			args: args{str: "abbaa"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.str); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPalindrome(t *testing.T) {
	type args struct {
		str   string
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Case1",
			args: args{
				str:   "cabac",
				left:  1,
				right: 3,
			},
			want: []string{"aba", "cabac"},
		},
		{
			name: "Case2",
			args: args{
				str:   "cabbac",
				left:  2,
				right: 3,
			},
			want: []string{"bb", "abba", "cabbac"},
		},
		{
			name: "Case3",
			args: args{
				str:   "test",
				left:  1,
				right: 2,
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPalindrome(tt.args.str, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findAllPalindrome(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Case1",
			args: args{"cabac"},
			want: []string{"aba", "cabac"},
		},
		{
			name: "Case2",
			args: args{"ouiuoabcdtat"},
			want: []string{"uiu", "ouiuo", "tat"},
		},
		{
			name: "Case3",
			args: args{""},
			want: []string{},
		},
		{
			name: "Case4",
			args: args{"b"},
			want: []string{"b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findAllPalindrome(tt.args.str)
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}
