package container

import (
	"strconv"
	"strings"
)

type node struct {
	data int
	next *node
}

func (n *node) Data() int {
	return n.data
}

func (n *node) SetNextNode(next *node) {
	n.next = next
}

func newNode(data int) *node {
	return &node{
		data: data,
		next: nil,
	}
}

type SinglyLinkedList struct {
	head *node
}

func NewSinglyLinkedList(ints []int) *SinglyLinkedList {
	if len(ints) <= 0 {
		return &SinglyLinkedList{
			head: nil,
		}
	}

	iterator := newNode(ints[0])
	list := &SinglyLinkedList{
		head: iterator,
	}
	for i := 1; i < len(ints); i++ {
		nextNode := newNode(ints[i])
		iterator.next = nextNode
		iterator = iterator.next
	}

	return list
}

func (s *SinglyLinkedList) Append(data int) {
	if s.head == nil {
		s.head = newNode(data)
		return
	}

	iterator := s.head
	for iterator.next != nil {
		iterator = iterator.next
	}

	iterator.next = newNode(data)
}

func (s *SinglyLinkedList) Preappend(data int) {
	newNode := newNode(data)
	newNode.next = s.head
	s.head = newNode
}

type NodeIndex = int

func (s *SinglyLinkedList) At(nodeIndex NodeIndex) *node {
	iterator := s.head
	for i := 0; i < nodeIndex; i++ {
		if iterator == nil {
			return nil
		}
		iterator = iterator.next
	}

	return iterator
}

func (s *SinglyLinkedList) FindNode(data int) NodeIndex {
	iterator := s.head
	var count int
	for iterator != nil {
		if iterator.data == data {
			return count
		}
		iterator = iterator.next
		count++
	}

	return -1
}

func (s *SinglyLinkedList) PopFront() {
	if s.head == nil {
		return
	}
	s.head = s.head.next
}

func (s *SinglyLinkedList) Remove(targetIndex NodeIndex) {
	if s.head == nil {
		return
	}

	if targetIndex == 0 {
		s.PopFront()
		return
	}

	iterator := s.head.next
	for i := 1; i < targetIndex-1; i++ {
		if iterator == nil {
			return
		}

		iterator = iterator.next
	}

	if iterator.next == nil {
		return
	} else if iterator.next.next == nil {
		iterator.next = nil
	} else {
		iterator.next = iterator.next.next
	}
}

func (s *SinglyLinkedList) GetSlice() []int {
	if s.head == nil {
		return []int{}
	}

	iterator := s.head
	var ints []int

	for iterator != nil {
		ints = append(ints, iterator.data)
		iterator = iterator.next
	}

	return ints
}

func (s *SinglyLinkedList) String() string {
	var chars []string

	iterator := s.head
	for iterator != nil {
		chars = append(chars, strconv.Itoa(iterator.data))
		iterator = iterator.next
	}

	return strings.Join(chars, " ")
}
