package main

import (
	"bytes"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "case1",
			input: "3\n5 1 2",
			want:  "8",
		},
		{
			name:  "case2",
			input: "5\n31 41 59 26 53",
			want:  "176",
		},
		{
			name:  "case3",
			input: "2\n31 22",
			want:  "9",
		},
		{
			name:  "case4",
			input: "2\n31 22",
			want:  "9",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// set up reader and writer
			input := &bytes.Buffer{}
			_, err := input.WriteString(test.input)
			if err != nil {
				t.Fatalf("failed to setup input data, err: %v", err)
			}
			reader = input

			result := &bytes.Buffer{}
			writer = result

			// execution
			main()

			// assertion
			if test.want != result.String() {
				t.Errorf("want: %v, but got: %v", test.want, result.String())
			}
		})
	}
}
