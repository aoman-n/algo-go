package quiz

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeleteDuplicateV1(t *testing.T) {
	type args struct {
		numbers *[]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case1",
			args: args{
				numbers: &[]int{1, 3, 3, 5, 5, 7, 7, 7, 10, 12, 12, 15},
			},
			want: []int{1, 3, 5, 7, 10, 12, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteDuplicateV1(tt.args.numbers)
			fmt.Printf("ret: %v\n", tt.args.numbers)
			if !reflect.DeepEqual(*tt.args.numbers, tt.want) {
				t.Errorf("got = %v, want %v", *tt.args.numbers, tt.want)
			}
		})
	}
}

func TestDeleteDuplicateV2(t *testing.T) {
	type args struct {
		numbers *[]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case1",
			args: args{
				numbers: &[]int{1, 3, 3, 5, 5, 7, 7, 7, 10, 12, 12, 15},
			},
			want: []int{1, 3, 5, 7, 10, 12, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteDuplicateV2(tt.args.numbers)
			if !reflect.DeepEqual(*tt.args.numbers, tt.want) {
				t.Errorf("got = %v, want %v", *tt.args.numbers, tt.want)
			}
		})
	}
}

func TestDeleteDuplicateV3(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case1",
			args: args{
				numbers: []int{1, 3, 3, 5, 5, 7, 7, 7, 10, 12, 12, 15},
			},
			want: []int{1, 3, 5, 7, 10, 12, 15},
		},
		{
			name: "Case2",
			args: args{
				numbers: []int{1, 3, 3, 5, 6, 6},
			},
			want: []int{1, 3, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteDuplicateV3(&tt.args.numbers)
			if !reflect.DeepEqual(tt.args.numbers, tt.want) {
				t.Errorf("got = %v, want %v", tt.args.numbers, tt.want)
			}
		})
	}
}

func TestDeleteDuplicateV4(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case1",
			args: args{
				numbers: []int{1, 3, 3, 5, 5, 7, 7, 7, 10, 12, 12, 15},
			},
			want: []int{1, 3, 5, 7, 10, 12, 15},
		},
		{
			name: "Case2",
			args: args{
				numbers: []int{1, 3, 3, 5, 6, 6},
			},
			want: []int{1, 3, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteDuplicateV4(&tt.args.numbers)
			if !reflect.DeepEqual(tt.args.numbers, tt.want) {
				t.Errorf("got = %v, want %v", tt.args.numbers, tt.want)
			}
		})
	}
}
