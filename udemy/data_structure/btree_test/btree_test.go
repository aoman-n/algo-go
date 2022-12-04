package btreetest

import (
	"fmt"
	"os"
	"testing"
)

func TestBtree(t *testing.T) {
	fmt.Println("hello!!!!")
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
	tree.Insert(20)
	tree.Insert(25)
	tree.Insert(30)
	tree.Insert(35)

	tree.InOrderPrint(os.Stdout)

	fmt.Printf("ret=%v\n", tree.Search(5))
	fmt.Printf("ret=%v\n", tree.Search(10))
	fmt.Printf("ret=%v\n", tree.Search(15))
	fmt.Printf("ret=%v\n", tree.Search(9))
	fmt.Printf("ret=%v\n", tree.Search(25))
	fmt.Printf("ret=%v\n", tree.Search(30))

	// n := tree.find(5)
	// if *n != nil {
	// 	fmt.Printf("n: %v\n", (**n).val)
	// }

	// n2 := tree.find(8)
	// if *n2 != nil {
	// 	fmt.Printf("n2: %v\n", (**n2).val)
	// }
}
