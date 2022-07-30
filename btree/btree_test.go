package btree

import (
	"fmt"
	"testing"
)

func TestBtree(t *testing.T) {
	tree := Tree[int]{
		cmp: func(a, b int) int {
			return a - b
		},
		root: &node[int]{
			left:  nil,
			right: nil,
			val:   10,
		},
	}

	tree.Insert(5)
	tree.Insert(15)

	n := tree.find(5)
	if *n != nil {
		fmt.Printf("n: %v\n", (**n).val)
	}

	n2 := tree.find(8)
	if *n2 != nil {
		fmt.Printf("n2: %v\n", (**n2).val)
	}
}
