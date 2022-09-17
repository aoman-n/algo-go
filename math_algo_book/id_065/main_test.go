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
			input: "3 3",
			want:  "5",
		},
		{
			name:  "case2",
			input: "4 5",
			want:  "10",
		},
		{
			name:  "case3",
			input: "7 3",
			want:  "11",
		},
		{
			name:  "case4",
			input: "1000000000 1000000000",
			want:  "500000000000000000",
		},
		{
			name:  "case5",
			input: "1 1",
			want:  "1",
		},
		{
			name:  "case6",
			input: "1 0",
			want:  "0",
		},
		{
			name:  "case7",
			input: "1 1000",
			want:  "1",
		},
		{
			name:  "case8",
			input: "1000 1",
			want:  "1",
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
