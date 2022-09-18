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
			input: "100 3\n2 3 5",
			want:  "74",
		},
		{
			name:  "case2",
			input: "100 3\n2 4 6",
			want:  "50",
		},
		{
			name:  "case3",
			input: "10000000000000 10\n13 17 19 23 29 31 37 41 43 47",
			want:  "3324865541894",
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
