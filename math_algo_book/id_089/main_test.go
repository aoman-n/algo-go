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
			input: "4 3 2",
			want:  "Yes",
		},
		{
			name:  "case2",
			input: "16 3 2",
			want:  "No",
		},
		{
			name:  "case3",
			input: "8 3 2",
			want:  "No",
		},
		{
			name:  "case4",
			input: "1000000000000000000 1000000000000000000 1000000000000000000",
			want:  "Yes",
		},
		{
			name:  "case5",
			input: "869120 5 15",
			want:  "No",
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
