package stack

import (
	"fmt"
	"io"
)

type Optional[T any] struct {
	Valid bool
	Val   T
}

type node[T any] struct {
	val  T
	next *node[T]
}

type Stack[T any] struct {
	head *node[T]
	len  int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		head: nil,
		len:  0,
	}
}

func (s *Stack[T]) Push(val T) {
	s.len++
	n := &node[T]{val, nil}
	if s.head == nil {
		s.head = n
		return
	}

	n.next = s.head
	s.head = n
}

func (s *Stack[T]) Peek() Optional[T] {
	if s.head == nil {
		return Optional[T]{}
	}

	return Optional[T]{Valid: true, Val: s.head.val}
}

func (s *Stack[T]) Pop() Optional[T] {
	if s.head == nil {
		return Optional[T]{}
	}

	tmp := s.head
	s.head = s.head.next
	s.len--
	return Optional[T]{
		Valid: true,
		Val:   tmp.val,
	}
}

func (s *Stack[T]) Print(w io.Writer) {
	first := true
	for n := s.head; n != nil; n = n.next {
		if first {
			fmt.Fprintf(w, "%v", n.val)
			first = false
		} else {
			fmt.Fprintf(w, " %v", n.val)
		}
	}
	fmt.Fprintln(w)
}

func (s *Stack[T]) Len() int {
	return s.len
}
