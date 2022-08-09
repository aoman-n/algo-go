package quiz

import "testing"

func TestListToIntPlusOne(t *testing.T) {
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
			args: args{[]int{1}},
			want: 2,
		},
		{
			name: "Case2",
			args: args{[]int{2, 3}},
			want: 24,
		},
		{
			name: "Case3",
			args: args{[]int{8, 9}},
			want: 90,
		},
		{
			name: "Case4",
			args: args{[]int{9, 9}},
			want: 100,
		},
		{
			name: "Case5",
			args: args{[]int{1, 2, 3}},
			want: 124,
		},
		{
			name: "Case6",
			args: args{[]int{9, 9, 9, 9}},
			want: 10000,
		},
		{
			name: "Case6",
			args: args{[]int{0, 0, 0, 9, 9, 9, 9}},
			want: 10000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListToIntPlusOne(tt.args.numbers); got != tt.want {
				t.Errorf("ListToIntPlusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
