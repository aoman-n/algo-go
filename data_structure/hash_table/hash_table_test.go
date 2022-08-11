package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashTable(t *testing.T) {
	h := NewHashTable()
	h.Add("car", "Tesla")
	h.Add("car", "Toyota")
	h.Add("pc", "Mac")
	h.Add("sns", "Youtube")
	h.Add("abc", "ok")

	assert.Equal(t, "Toyota", h.Get("car"))
	assert.Equal(t, "Mac", h.Get("pc"))
	assert.Equal(t, "Youtube", h.Get("sns"))
	assert.Equal(t, "ok", h.Get("abc"))
}
