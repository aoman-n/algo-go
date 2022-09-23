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
			input: "2\n3 5",
			want:  "13",
		},
		{
			name:  "case2",
			input: "3\n1 3 5",
			want:  "27",
		},
		{
			name:  "case3",
			input: "1\n5",
			want:  "5",
		},
		{
			name:  "case4",
			input: "2\n11 11",
			want:  "33",
		},
		{
			name:  "case5",
			input: "2\n11 11",
			want:  "33",
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
