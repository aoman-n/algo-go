package main

import (
	"bytes"
	"fmt"

	//lint:ignore SA1019 use Go 1.14.1 at AtCoder
	"io/ioutil"
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	tests, err := ioutil.ReadDir("./testfiles")
	if err != nil {
		t.Fatal(err)
	}
	for i, test := range tests {
		if i != 1 {
			continue
		}
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

// 添字
//  0  1  2  3  4  5
// アクセス(normal)
//  0  1  2  3  4  5
// アクセス(reverse)
//  6  5  4  3  2  1  0
// [0, 1, 2, 3, 4, 5, ?]

// len(7)要素数(5)
// 7(len)-1(x)-1=5
