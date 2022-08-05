package quiz

import "testing"

func Test_countChars(t *testing.T) {
	type args struct {
		chars string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		{
			name:  "Case1",
			args:  args{"This is a pen. This is an apple. Applepen."},
			want:  "p",
			want1: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := countChars(tt.args.chars)
			if got != tt.want {
				t.Errorf("countChars() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("countChars() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
