package main

import (
	"reflect"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case1",
			args: args{list: []int{8, 3, 4, 9, 1, 2}},
			want: []int{1, 2, 3, 4, 8, 9},
		},
		{
			name: "Case1",
			args: args{list: []int{681, 857, 76, 483, 602, 586}},
			want: []int{76, 483, 586, 602, 681, 857},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergeSort(tt.args.list)
			if !reflect.DeepEqual(tt.args.list, tt.want) {
				t.Errorf("mergeSort() = %v, but want %v", tt.args.list, tt.want)
			}
		})
	}
}
