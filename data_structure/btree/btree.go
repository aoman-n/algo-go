package btree

import (
	"fmt"
	"io"
)

type node[T any] struct {
	val         T
	left, right *node[T]
}

func newNode[T any](val T) *node[T] {
	return &node[T]{val, nil, nil}
}

type BST[T any] struct {
	cmp  func(T, T) int
	root *node[T]
}

func NewBST[T any](cmp func(T, T) int) *BST[T] {
	return &BST[T]{
		cmp:  cmp,
		root: nil,
	}
}

func (b *BST[T]) Insert(val T) {
	if b.root == nil {
		b.root = &node[T]{val, nil, nil}
		return
	}

	b.insert(b.root, val)
}

func (b *BST[T]) insert(root *node[T], val T) *node[T] {
	if root == nil {
		return newNode(val)
	}

	cmp := b.cmp(root.val, val)
	if cmp > 0 {
		root.left = b.insert(root.left, val)
	} else if cmp < 0 {
		root.right = b.insert(root.right, val)
	}

	return root
}

func (b *BST[T]) Insert2(val T) bool {
	n := b.find(val)

	if *n != nil {
		return false
	}

	*n = newNode(val)
	return true
}

func (b *BST[T]) find(val T) **node[T] {
	pl := &b.root

	for *pl != nil {
		switch cmp := b.cmp((*pl).val, val); {
		case cmp > 0:
			pl = &(*pl).left
		case cmp < 0:
			pl = &(*pl).right
		default:
			return pl
		}
	}

	return pl
}

func (b *BST[T]) InOrderPrint(w io.Writer) {
	b.inOrderPrint(w, b.root)
	fmt.Fprintln(w)
}

func (b *BST[T]) inOrderPrint(w io.Writer, root *node[T]) {
	if root != nil {
		b.inOrderPrint(w, root.left)
		fmt.Fprintf(w, "%v ", root.val)
		b.inOrderPrint(w, root.right)
	}
}

func (b *BST[T]) InOrderSlice() []T {
	s := make([]T, 0)
	b.inOrderSlice(b.root, &s)
	return s
}

func (b *BST[T]) inOrderSlice(node *node[T], s *[]T) {
	if node != nil {
		b.inOrderSlice(node.left, s)
		*s = append(*s, node.val)
		b.inOrderSlice(node.right, s)
	}
}

func (b *BST[T]) Search(val T) bool {
	return b.search(b.root, val)
}

func (b *BST[T]) search(root *node[T], val T) bool {
	if root == nil {
		return false
	}

	switch cmp := b.cmp(root.val, val); {
	case cmp > 0:
		return b.search(root.left, val)
	case cmp < 0:
		return b.search(root.right, val)
	default:
		return true
	}
}

func (b *BST[T]) Remove(val T) {
	b.root = b.remove(b.root, val)
}

func (b *BST[T]) remove(root *node[T], val T) *node[T] {
	if root == nil {
		return nil
	}

	switch cmp := b.cmp(root.val, val); {
	case cmp > 0:
		root.left = b.remove(root.left, val)
	case cmp < 0:
		root.right = b.remove(root.right, val)
	default:
		if root.right == nil {
			return root.left
		} else if root.left == nil {
			return root.right
		} else {
			tmp := b.minNode(root.right)
			root.val = tmp.val
			root.right = b.remove(root.right, tmp.val)
		}
	}

	return root
}

func (b *BST[T]) minNode(root *node[T]) *node[T] {
	currentNode := root
	for currentNode.left != nil {
		currentNode = currentNode.left
	}
	return currentNode
}
