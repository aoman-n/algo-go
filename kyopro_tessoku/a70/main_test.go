package main

import (
	"bytes"
	"fmt"

	"io/ioutil"
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	tests, err := ioutil.ReadDir("./testfiles")
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		t.Run(test.Name(), func(t *testing.T) {
			// set up reader and writer
			input, _ := os.Open(fmt.Sprintf("./testfiles/%s/input.txt", test.Name()))
			reader = input
			result := &bytes.Buffer{}
			writer = result

			// execution
			main()

			// assertion
			expected, _ := os.ReadFile(fmt.Sprintf("./testfiles/%s/output.txt", test.Name()))
			if result.String() != string(expected) {
				t.Logf("expected: %s\n", string(expected))
				t.Fatalf("unexpected result: %s\n", result.String())
			}
		})
	}
}

func TestOperation_nextVIndex(t *testing.T) {
	tests := []struct {
		operation *operation
		input     int
		want      int
	}{
		{
			operation: newOperation(1, 2, 3),
			input:     0,
			want:      7,
		},
		{
			operation: newOperation(1, 2, 3),
			input:     7,
			want:      0,
		},
		{
			operation: newOperation(2, 3, 4),
			input:     0,
			want:      14,
		},
		{
			operation: newOperation(2, 3, 4),
			input:     14,
			want:      0,
		},
		{
			operation: newOperation(2, 3, 4),
			input:     9,
			want:      7,
		},
		{
			operation: newOperation(2, 3, 4),
			input:     10,
			want:      4,
		},
	}

	for i, tt := range tests {
		tt := tt

		t.Run(fmt.Sprintf("case(%v)", i+1), func(t *testing.T) {
			got := tt.operation.nextVIndex(tt.input, 10)
			if got != tt.want {
				t.Errorf("operation.nextVIndex(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
