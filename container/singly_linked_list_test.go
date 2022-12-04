package container_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/aoman-n/algo-go/container"
	"github.com/google/go-cmp/cmp"
)

func TestSinglyLinkedList_New(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input []int
		want  string
	}{
		"zero length data": {[]int{}, ""},
		"case1":            {[]int{10, 20, 30, 40, 50}, "10 20 30 40 50"},
		"case2":            {[]int{-1, 20, -10, 3}, "-1 20 -10 3"},
	}

	for name, tt := range tests {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			list := container.NewSinglyLinkedList(tt.input)

			buf := &bytes.Buffer{}
			fmt.Fprint(buf, list)
			got := buf.String()

			if got != tt.want {
				t.Errorf("NewSinglyLinkedList(%v) = %v, got %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestSinglyLinkedList_Append_Preappend_PopFront(t *testing.T) {
	list := &container.SinglyLinkedList{}

	// Test Append
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)
	list.Append(50)
	buf := &bytes.Buffer{}
	fmt.Fprint(buf, list)

	got := buf.String()
	want := "10 20 30 40 50"
	if got != want {
		t.Errorf("SinglyLinkedList.Append() = %v, want %v", got, want)
	}

	// Test Preappend
	list.Preappend(0)
	list.Preappend(-10)
	list.Preappend(-20)
	list.Preappend(-30)
	list.Preappend(-40)

	buf2 := &bytes.Buffer{}
	fmt.Fprint(buf2, list)
	got2 := buf2.String()
	want2 := "-40 -30 -20 -10 0 10 20 30 40 50"
	if got2 != want2 {
		t.Errorf("SinglyLinkedList.Preappend() = %v, want %v", got2, want2)
	}

	// Test PopFront
	list.PopFront()
	list.PopFront()
	list.PopFront()

	buf3 := &bytes.Buffer{}
	fmt.Fprint(buf3, list)
	got3 := buf3.String()
	want3 := "-10 0 10 20 30 40 50"
	if got3 != want3 {
		t.Errorf("SinglyLinkedList.Preappend() = %v, want %v", got3, want3)
	}
}

func TestSinglyLinkedList_Remove(t *testing.T) {
	list := container.NewSinglyLinkedList([]int{10, 20, 30, 40, 50})

	list.Remove(2)
	got := list.GetSlice()
	want := []int{10, 20, 40, 50}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Fatalf("Remove(2) mismatch (-want +got):\n%s", diff)
	}

	list.Remove(0)
	got1 := list.GetSlice()
	want1 := []int{20, 40, 50}
	if diff := cmp.Diff(got1, want1); diff != "" {
		t.Fatalf("Remove(0) mismatch (-want +got):\n%s", diff)
	}

	list.Remove(2)
	got2 := list.GetSlice()
	want2 := []int{20, 40}
	if diff := cmp.Diff(got2, want2); diff != "" {
		t.Fatalf("Remove(2) mismatch (-want +got):\n%s", diff)
	}

	list.Remove(2)
	got3 := list.GetSlice()
	want3 := []int{20, 40}
	if diff := cmp.Diff(got3, want3); diff != "" {
		t.Fatalf("Remove(2) mismatch (-want +got):\n%s", diff)
	}
}

func TestSinglyLinkedList_At(t *testing.T) {
	list := &container.SinglyLinkedList{}
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)
	list.Append(50)

	node1 := list.At(3)
	want1 := 40
	if node1.Data() != want1 {
		t.Errorf("node.At(3) = %v, want %v", node1.Data(), want1)
	}

	node2 := list.At(0)
	want2 := 10
	if node2.Data() != want2 {
		t.Errorf("node.At(0) = %v, want %v", node2.Data(), want2)
	}
}

func TestSinglyLinkedList(t *testing.T) {
	t.Parallel()

	list := &container.SinglyLinkedList{}
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)
	list.Append(50)

	tests := map[string]struct {
		input int
		want  int
	}{
		"[FindNode] exists case 1":    {10, 0},
		"[FindNode] exists case 2":    {50, 4},
		"[FindNode] not found case 1": {222, -1},
	}

	for name, tt := range tests {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := list.FindNode(tt.input)

			if got != tt.want {
				t.Errorf("node.FindNode(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
