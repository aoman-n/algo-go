package btree

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBST(t *testing.T) {
	fmt.Println("hello world!!")

	b := NewBST(func(i1, i2 int) int { return i1 - i2 })

	b.Insert2(3)
	b.Insert2(6)
	b.Insert2(5)
	b.Insert2(7)
	b.Insert2(1)
	b.Insert2(10)
	b.Insert2(2)
	b.InOrderPrint(os.Stdout)

	assert.Equal(t, true, b.Search(3))
	assert.Equal(t, true, b.Search(5))
	assert.Equal(t, false, b.Search(8))
	assert.Equal(t, false, b.Search(11))

	b.Remove(11)
	assert.Equal(t, []int{1, 2, 3, 5, 6, 7, 10}, b.InOrderSlice())
	b.Remove(6)
	assert.Equal(t, []int{1, 2, 3, 5, 7, 10}, b.InOrderSlice())
	b.Remove(10)
	assert.Equal(t, []int{1, 2, 3, 5, 7}, b.InOrderSlice())
	b.Remove(1)
	assert.Equal(t, []int{2, 3, 5, 7}, b.InOrderSlice())
	b.Remove(5)
	assert.Equal(t, []int{2, 3, 7}, b.InOrderSlice())
}
