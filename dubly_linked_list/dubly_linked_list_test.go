package dllist

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	l := NewList()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)

	l.Insert(0)
	l.Insert(-1)
	l.Print(os.Stdout)

	l.Remove(-1)
	l.Print(os.Stdout)

	l.Remove(4)
	l.Print(os.Stdout)

	l.Remove(2)
	l.Print(os.Stdout)
}

func TestReverseIterative(t *testing.T) {
	tests := []struct {
		name    string
		prepare func() *List
		want    string
	}{
		{
			name: "Case1",
			prepare: func() *List {
				l := NewList()
				l.Append(1)
				l.Append(2)
				l.Append(3)
				l.Append(4)
				l.Append(5)
				return l
			},
			want: "5 4 3 2 1\n",
		},
		{
			name: "Case2",
			prepare: func() *List {
				l := NewList()
				l.Append(111)
				return l
			},
			want: "111\n",
		},
		{
			name: "Case3",
			prepare: func() *List {
				l := NewList()
				return l
			},
			want: "\n",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.Buffer{}
			l := tt.prepare()
			l.ReverseIterative()
			l.Print(&buf)

			got := buf.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want %v, but got %v", tt.want, got)
			}
		})
	}
}

func TestReverseRecursive(t *testing.T) {
	tests := []struct {
		name    string
		prepare func() *List
		want    string
	}{
		{
			name: "Case1",
			prepare: func() *List {
				l := NewList()
				l.Append(1)
				l.Append(2)
				l.Append(3)
				l.Append(4)
				l.Append(5)
				return l
			},
			want: "5 4 3 2 1\n",
		},
		{
			name: "Case2",
			prepare: func() *List {
				l := NewList()
				l.Append(111)
				return l
			},
			want: "111\n",
		},
		{
			name: "Case3",
			prepare: func() *List {
				l := NewList()
				return l
			},
			want: "\n",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.Buffer{}
			l := tt.prepare()
			l.ReverseRecursive()
			l.Print(&buf)

			got := buf.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want %v, but got %v", tt.want, got)
			}
		})
	}
}
