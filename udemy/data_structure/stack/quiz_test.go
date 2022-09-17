package stack

import "testing"

func TestValidateFormat(t *testing.T) {
	type args struct {
		chars string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case1",
			args: args{"{'key1': 'value1', 'key2': [1, 2, 3], 'key3': (1, 2, 3)}"},
			want: true,
		},
		{
			name: "Case2",
			args: args{"[{'key1': 'value1', 'key2': [1, 2, 3], 'key3': (1, 2, 3)}"},
			want: false,
		},
		{
			name: "Case3",
			args: args{"{'key1': 'value1', 'key2': [1, 2, 3], 'key3': (1, 2, 3)}]"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateFormat(tt.args.chars); got != tt.want {
				t.Errorf("ValidateFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
