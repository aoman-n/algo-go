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
			input: "1 2 3",
			want:  "18",
		},
		{
			name:  "case2",
			input: "1000000000 987654321 123456789",
			want:  "951633476",
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
