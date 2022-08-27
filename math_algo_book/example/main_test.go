package main

import (
	"bytes"
	"testing"
)

func Test_run(t *testing.T) {
	t.Run("10を入力するとChildが出力されること", func(t *testing.T) {
		reader := &bytes.Buffer{}
		reader.Write([]byte("10"))
		writer := &bytes.Buffer{}

		run(reader, writer)

		got := writer.String()
		if got != Child+"\n" {
			t.Errorf("want %v, but got %v", Child, got)
		}
	})

	t.Run("2を入力するとAdultが出力されること", func(t *testing.T) {
		reader := &bytes.Buffer{}
		reader.Write([]byte("20"))
		writer := &bytes.Buffer{}

		run(reader, writer)

		got := writer.String()
		if got != Adult+"\n" {
			t.Errorf("want %v, but got %v", Adult, got)
		}
	})
}
