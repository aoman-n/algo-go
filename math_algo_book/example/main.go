package main

import (
	"fmt"
	"io"
	"os"
)

const (
	Child = "Child"
	Adult = "Adult"
)

// @see: https://qiita.com/syumai/items/d4d436eacc58ffbd8200

func run(reader io.Reader, writer io.Writer) {
	var n int
	fmt.Fscan(reader, &n)

	if n < 20 {
		fmt.Fprintln(writer, Child)
		return
	}

	fmt.Fprintln(writer, Adult)
}

func main() {
	run(os.Stdin, os.Stdout)
}
