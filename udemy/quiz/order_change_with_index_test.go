package quiz

import "testing"

func Test_orderChangeByIndexesV1(t *testing.T) {
	type args struct {
		chars   []string
		indexes []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case1",
			args: args{
				chars:   []string{"h", "y", "n", "p", "t", "o"},
				indexes: []int{3, 1, 4, 0, 5, 2},
			},
			want: "python",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderChangeByIndexesV1(tt.args.chars, tt.args.indexes); got != tt.want {
				t.Errorf("orderChangeByIndexesV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderChangeByIndexesV2(t *testing.T) {
	type args struct {
		chars   []string
		indexes []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case1",
			args: args{
				chars:   []string{"h", "y", "n", "p", "t", "o"},
				indexes: []int{3, 1, 4, 0, 5, 2},
			},
			want: "python",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderChangeByIndexesV2(tt.args.chars, tt.args.indexes); got != tt.want {
				t.Errorf("orderChangeByIndexesV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
