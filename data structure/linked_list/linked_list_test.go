package linkedlist

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_Run(t *testing.T) {
}

func Test_ReverseEven(t *testing.T) {
	tests := []struct {
		name    string
		prepare func() *LinkedList
		want    string
	}{
		{
			name: "Case1",
			prepare: func() *LinkedList {
				l := NewLinkedList()
				l.Append(1)
				l.Append(4)
				l.Append(6)
				l.Append(8)
				l.Append(9)
				return l
			},
			want: "1 8 6 4 9 ",
		},
		{
			name: "Case2",
			prepare: func() *LinkedList {
				l := NewLinkedList()
				l.Append(1)
				l.Append(2)
				l.Append(3)
				l.Append(4)
				l.Append(5)
				l.Append(6)
				return l
			},
			want: "1 2 3 4 5 6 ",
		},
		{
			name: "Case2",
			prepare: func() *LinkedList {
				l := NewLinkedList()
				l.Append(1)
				l.Append(2)
				l.Append(4)
				l.Append(6)
				l.Append(8)
				l.Append(9)
				l.Append(10)
				l.Append(14)
				l.Append(12)
				return l
			},
			want: "1 8 6 4 2 9 12 14 10 ",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			l := tt.prepare()
			l.ReverseEven()

			buf := bytes.Buffer{}
			l.Print(&buf)
			if got := buf.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want %v, but got %v", tt.want, got)
			}
		})
	}
}
