package btreetest

import (
	"fmt"
	"io"
)

// Tree is a binary tree.
type Tree[T any] struct {
	cmp  func(T, T) int
	root *node[T]
}

// A node in a Tree.
type node[T any] struct {
	left, right *node[T]
	val         T
}

func (b *Tree[T]) InOrderPrint(w io.Writer) {
	b.inOrderPrint(w, b.root)
	fmt.Fprintln(w)
}

func (b *Tree[T]) inOrderPrint(w io.Writer, root *node[T]) {
	if root != nil {
		b.inOrderPrint(w, root.left)
		fmt.Fprintf(w, "%v ", root.val)
		b.inOrderPrint(w, root.right)
	}
}

// find returns a pointer to the node containing val,
// or, if val is not present, a pointer to where it
// would be placed if added.
// func (bt *Tree[T]) find(val T) **node[T] {
// 	pl := &bt.root
// 	for *pl != nil {
// 		switch cmp := bt.cmp(val, (*pl).val); {
// 		case cmp < 0:
// 			pl = &(*pl).left
// 		case cmp > 0:
// 			pl = &(*pl).right
// 		default:
// 			return pl
// 		}
// 	}
// 	return pl
// }

// Insert inserts val into bt if not already there,
// and reports whether it was inserted.
// func (bt *Tree[T]) Insert(val T) bool {
// 	pl := bt.find(val)
// 	if *pl != nil {
// 		return false
// 	}
// 	*pl = &node[T]{val: val}
// 	return true
// }

func newNode[T any](val T) *node[T] {
	return &node[T]{
		left:  nil,
		right: nil,
		val:   val,
	}
}

func (bt *Tree[T]) Insert(val T) bool {
	pl := bt.root

outer:
	for pl != nil {
		switch cmp := bt.cmp(val, pl.val); {
		case cmp < 0:
			if pl.left == nil {
				pl.left = newNode(val)
				break outer
			}
			pl = pl.left
		case cmp > 0:
			if pl.right == nil {
				pl.right = newNode(val)
				break outer
			}
			pl = pl.right
		default:
			return false
		}
	}

	return true
}

func (bt *Tree[T]) Search(val T) bool {
	pl := bt.root

	for pl != nil {
		switch cmp := bt.cmp(val, pl.val); {
		case cmp < 0:
			pl = pl.left
		case cmp > 0:
			pl = pl.right
		default:
			return true
		}
	}

	return false
}
